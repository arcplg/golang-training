<template>
    <div id="main-wrapper" class="auth-customizer-none">
        <div class="position-relative overflow-hidden radial-gradient min-vh-100 w-100">
            <div class="position-relative z-index-5">
                <div class="row">
                    <div class="col-xl-7 col-xxl-8">
                        <a href="#" class="text-nowrap logo-img d-block px-4 py-9 w-100">
                            <img src="/assets/images/logos/dark-logo.svg" class="dark-logo" alt="Logo-Dark">
                            <!-- <img src="/assets/images/logos/light-logo.svg" class="light-logo" alt="Logo-light"> -->
                        </a>
                        <div class="d-none d-xl-flex align-items-center justify-content-center h-n80">
                            <img src="/assets/images/backgrounds/login-security.svg" alt="modernize-img" class="img-fluid" width="500">
                        </div>
                    </div>
                    <div class="col-xl-5 col-xxl-4">
                        <div class="authentication-login min-vh-100 bg-body row justify-content-center align-items-center p-4">
                            <div class="auth-max-width col-sm-8 col-md-6 col-xl-7 px-4">
                                <h2 class="mb-1 fs-7 fw-bolder">Welcome to Modernize</h2>
                                <p class="mb-7">Your Admin Dashboard</p>
                                <div class="row">
                                    <div class="col-6 mb-2 mb-sm-0">
                                        <a class="btn text-dark border fw-normal d-flex align-items-center justify-content-center rounded-2 py-8" href="javascript:void(0)" role="button">
                                            <img src="/assets/images/svgs/google-icon.svg" alt="modernize-img" class="img-fluid me-2" width="18" height="18">
                                            <span class="flex-shrink-0">with Google</span>
                                        </a>
                                    </div>
                                    <div class="col-6">
                                        <a class="btn text-dark border fw-normal d-flex align-items-center justify-content-center rounded-2 py-8" href="javascript:void(0)" role="button">
                                            <img src="/assets/images/svgs/facebook-icon.svg" alt="modernize-img" class="img-fluid me-2" width="18" height="18">
                                            <span class="flex-shrink-0">with FB</span>
                                        </a>
                                    </div>
                                </div>
                                <div class="position-relative text-center my-4">
                                    <p class="mb-0 fs-4 px-3 d-inline-block bg-body text-dark z-index-5 position-relative">or sign in with</p>
                                    <span class="border-top w-100 position-absolute top-50 start-50 translate-middle"></span>
                                </div>
                                <form>
                                    <div class="mb-3">
                                        <label for="exampleInputUsername" class="form-label">Username</label>
                                        <input type="email" class="form-control" id="exampleInputUsername" aria-describedby="userNameHelp">
                                    </div>
                                    <div class="mb-3">
                                        <label for="exampleInputEmail1" class="form-label">Email</label>
                                        <input type="email" class="form-control" id="exampleInputEmail1" aria-describedby="emailHelp">
                                    </div>
                                    <div class="mb-4">
                                        <label for="exampleInputPassword1" class="form-label">Password</label>
                                        <input type="password" class="form-control" id="exampleInputPassword1">
                                    </div>
                                    <a href="#" class="btn btn-primary w-100 py-8 mb-4 rounded-2">Register account</a>
                                    <div class="d-flex align-items-center justify-content-center">
                                        <p class="fs-4 mb-0 fw-medium">You have account?</p>
                                        <a class="text-primary fw-medium ms-2" href="#">Sign In</a>
                                    </div>
                                </form>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <button class="btn btn-primary p-3 rounded-circle d-flex align-items-center justify-content-center customizer-btn" type="button" data-bs-toggle="offcanvas" data-bs-target="#offcanvasExample" aria-controls="offcanvasExample">
            <i class="icon ti ti-settings fs-7"></i>
        </button>
    </div>
</template>
<script>
    export default {
        data() {
            return {
                username: '',
                password: '',
                rememberMe: false
            };
        },
        methods: {
            login() {
                const payload = {
                    username: this.username,
                    password: this.password,
                    rememberMe: this.rememberMe
                };
                // Implement login logic (send API request, etc.)
                console.log(payload);
            }
        }
    };
</script>
<style scoped>
    /* Add any specific styles if needed */
</style>
