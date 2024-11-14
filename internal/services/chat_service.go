package services

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/ngocthanh06/chatapp/internal/handlers"
	"github.com/ngocthanh06/chatapp/internal/middleware"
	"github.com/ngocthanh06/chatapp/internal/models"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

type ChatAppServiceServer struct {
	handlers.UnimplementedChatAppServiceServer
	Db      *gorm.DB
	mu      sync.Mutex
	Clients map[handlers.ChatAppService_ChatWithUserStreamServer]string
	MsgChan chan *handlers.ChatWithUserMessage
}

func (c *ChatAppServiceServer) LoginUser(ctx context.Context, req *handlers.LoginRequest) (*handlers.LoginResponse, error) {
	log.Println("Login User is running...!")

	// Check validation
	if req.Username == "" {
		return nil, status.Errorf(codes.InvalidArgument, "request missing required field: Username")
	} else if req.Password == "" {
		return nil, status.Errorf(codes.InvalidArgument, "request missing required field: Password")
	}

	var user models.User
	err := c.Db.Where("username = ?", req.Username).First(&user).Error

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	// check has password
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, status.Errorf(codes.NotFound, "password is not correct")
	}

	//Create claims with multiple fields populated
	claims := middleware.Claims{
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			ID:        user.Id.String(),
		},
		user.Username,
		user.Id.String(),
	}

	// Create a new token object, specifying signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))

	return &handlers.LoginResponse{Token: tokenStr}, nil
}

func (c *ChatAppServiceServer) RegisterUser(ctx context.Context, req *handlers.RegisterRequest) (*handlers.RegisterResponse, error) {
	log.Println("Register User is running...!")

	// Check validation
	if req.Username == "" {
		return nil, status.Errorf(codes.InvalidArgument, "request missing required field: Username")
	} else if req.Password == "" {
		return nil, status.Errorf(codes.InvalidArgument, "request missing required field: Password")
	}

	// check user exist, if no then create new
	var user models.User
	result := c.Db.Where("username = ?", req.Username).First(&user)

	if result.Error == nil {
		return nil, status.Errorf(codes.AlreadyExists, "user already exist")
	} else if result.Error != gorm.ErrRecordNotFound {
		return nil, status.Errorf(codes.Internal, "database error %v", result.Error)
	}

	var hashPassword []byte
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Generate password is wrong!")
	}

	user = models.User{
		Username: req.Username,
		Password: string(hashPassword),
	}

	if err := c.Db.Create(&user).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user: %v", err)
	}

	return &handlers.RegisterResponse{
		UserId: user.Id.String(),
	}, nil
}

func (c *ChatAppServiceServer) CreateRoom(ctx context.Context, req *handlers.CreateRoomRequest) (*handlers.CreateRoomResponse, error) {
	log.Println("Create room is running...!")

	if nameRoom := req.RoomName; nameRoom == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Room name is required")
	}

	var room models.Room

	result := c.Db.Where("name = ?", req.RoomName).First(&room)

	if result.Error == nil {
		return nil, status.Errorf(codes.AlreadyExists, "room already exist")
	} else if result.Error != gorm.ErrRecordNotFound {
		return nil, status.Errorf(codes.Internal, "database error %v", result.Error)
	}

	room = models.Room{
		Name: req.RoomName,
	}

	if err := c.Db.Create(&room).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create room %v", err)
	}

	return &handlers.CreateRoomResponse{RoomId: room.Id.String()}, nil
}

func (c *ChatAppServiceServer) JoinRoom(ctx context.Context, req *handlers.JoinRoomRequest) (*handlers.JoinRoomResponse, error) {
	log.Println("Join Room is calling...!")

	// check validation
	if req.RoomId == "" || req.UserId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Room id and user id are required!!")
	}

	if _, err := uuid.Parse(req.RoomId); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid room Id is UUID")
	}

	if _, err := uuid.Parse(req.UserId); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid user Id is UUID")
	}

	var user models.User

	// check user exist
	resultUser := c.Db.Where("id", req.UserId).Find(&user)

	if resultUser.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "User not found")
	}

	var room models.Room
	// check room exist
	resultRoom := c.Db.Where("id", req.RoomId).Find(&room)

	if resultRoom.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Room not found")
	}

	var roomMember models.RoomMember

	result := c.Db.Where("user_id", req.UserId).Where("room_id", req.RoomId).First(&roomMember)

	if result.Error == nil {
		return nil, status.Errorf(codes.AlreadyExists, "Member is joining group.!")
	} else if result.Error != gorm.ErrRecordNotFound {
		return nil, status.Errorf(codes.Internal, "database error %v", result.Error)
	}

	roomId, _ := uuid.Parse(req.RoomId)
	userId, _ := uuid.Parse(req.UserId)

	roomMember = models.RoomMember{
		RoomId:   roomId,
		UserId:   userId,
		JoinedAt: time.Now(),
	}

	if err := c.Db.Create(&roomMember).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create room %v", err)
	}

	return &handlers.JoinRoomResponse{
		Success: true,
	}, nil
}

func (c *ChatAppServiceServer) ChatWithUserStream(stream grpc.BidiStreamingServer[handlers.ChatWithUserMessage, handlers.ChatWithUserMessage]) error {
	log.Println("Chat Stream calling...!")

	// insert client into map
	c.mu.Lock()
	c.Clients[stream] = ""
	c.mu.Unlock()

	defer func() {
		c.mu.Lock()
		delete(c.Clients, stream)
		c.mu.Unlock()
	}()

	errChan := make(chan error)

	go c.receiveMessages(stream, errChan)

	for {
		select {
		case err := <-errChan:
			if err != nil {
				return err
			}
		case msg := <-c.MsgChan:
			if err := c.sendMessageToClients(msg); err != nil {
				return err
			}
		}
	}
}

func (c *ChatAppServiceServer) sendMessageToClients(msg *handlers.ChatWithUserMessage) error {
	// send message
	c.mu.Lock()
	defer c.mu.Unlock()
	for client, receiver := range c.Clients {
		if msg.Receiver != receiver {
			continue
		}

		if err := client.Send(msg); err != nil {
			return status.Errorf(codes.Internal, "Send message error: %v", err)
		}
	}
	return nil
}

func (c *ChatAppServiceServer) receiveMessages(stream grpc.BidiStreamingServer[handlers.ChatWithUserMessage, handlers.ChatWithUserMessage], errChan chan error) {
	defer close(errChan)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			errChan <- nil
			return
		}

		if err != nil {
			errChan <- status.Errorf(codes.InvalidArgument, err.Error())
			return
		}

		if err = c.validateMessage(req); err != nil {
			errChan <- err
			return
		}

		c.mu.Lock()
		if _, exists := c.Clients[stream]; !exists {
			c.Clients[stream] = req.Sender
		} else if c.Clients[stream] == "" {
			c.Clients[stream] = req.Sender
		}
		c.mu.Unlock()

		c.MsgChan <- req
	}
}

func (c *ChatAppServiceServer) validateMessage(req *handlers.ChatWithUserMessage) error {
	if req == nil {
		return status.Errorf(codes.InvalidArgument, "Params is required")
	}

	if req.Receiver == "" {
		return status.Errorf(codes.InvalidArgument, "Params is required")
	}

	if req.Sender == "" {
		return status.Errorf(codes.InvalidArgument, "Params is required")
	}

	if req.Message == "" {
		return status.Errorf(codes.InvalidArgument, "Params is required")
	}

	// validation receiver is uuid
	if _, errReceiver := uuid.Parse(req.Receiver); errReceiver != nil {
		return status.Errorf(codes.InvalidArgument, "Receiver is not a valid UUID: %v\", err")
	}

	return nil
}
