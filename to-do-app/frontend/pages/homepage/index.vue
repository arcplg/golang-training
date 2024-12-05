
<style scoped>
    .form-control {
        padding: 0 !important;
    }
</style>
<template>
    <div class="flex w-full min-h-screen">
        <div class="page-wrapper flex w-full">
            <Sidebar />
            <div class="body-wrapper w-full bg-white dark:bg-dark">
                <div aria-modal="true" aria-describedby="drawer-dialog-:R2cvaja:" role="dialog" tabindex="-1" data-testid="flowbite-drawer" class="fixed z-40 overflow-y-auto bg-white dark:bg-dark p-0 transition-transform left-0 top-0 h-screen -translate-x-full w-130">
                    <div data-testid="flowbite-drawer-items" class="">
                        <div class="flex">
                            <nav aria-label="Sidebar with multi-level dropdown example" class="h-full w-64 fixed menu-sidebar pt-0 bg-white dark:bg-dark z-[10]">
                                <div class="bg-white dark:bg-dark rounded-none w-[270px] border-r-4 border-primary dark:border-darkborder h-full">
                                    <div class="px-6 flex items-center brand-logo overflow-hidden">
                                        <a href="/">
                                            <img alt="logo" loading="lazy" width="174" height="26" decoding="async" data-nimg="1" class="block dark:hidden rtl:scale-x-[-1]" style="color:transparent" src="/assets/images/dark-logo.31875ba3.svg">
                                            <img alt="logo" loading="lazy" width="174" height="26" decoding="async" data-nimg="1" class="hidden dark:block rtl:scale-x-[-1]" style="color:transparent" src="/assets/images/light-logo.48dc618e.svg">
                                        </a>
                                    </div>
                                </div>
                            </nav>
                        </div>
                    </div>
                </div>
                <div class=" container mx-auto  p-6 ">
                    <div data-testid="flowbite-card" class="flex relative w-full break-words flex-col card p-6 dark:shadow-dark-md mb-6 py-4 bg-lightinfo dark:bg-darkinfo overflow-hidden rounded-md border-none shadow-none dark:shadow-none" style="border-radius:7px">
                        <div class="flex h-full flex-col justify-start gap-0 p-0">
                            <div class=" items-center grid grid-cols-12 gap-6">
                                <div class="col-span-9">
                                    <h4 class="font-semibold text-xl text-dark dark:text-white mb-3">To do list</h4>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div data-testid="flowbite-card" class="flex rounded-md bg-white dark:bg-dark relative w-full break-words flex-col card p-6 dark:shadow-dark-md shadow-md undefined" style="border-radius:7px">
                        <div class="flex h-full flex-col justify-start gap-0 p-0">
                            <div class="flex gap-3 justify-between items-center mb-5">
                                <div class="flex form-control w-full sm:max-w-60 max-w-full">
                                    <div class="relative w-full">
                                        <div class="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3">
                                            <span></span>
                                        </div>
                                        <input class="block w-full border disabled:cursor-not-allowed disabled:opacity-50 border-gray-300 bg-gray-50 text-gray-900 focus:border-cyan-500 focus:ring-cyan-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-cyan-500 dark:focus:ring-cyan-500 p-2.5 text-sm pl-10 rounded-lg" type="text" id="search" placeholder="Search Products" required="" value="">
                                    </div>
                                </div>
                                <div class="flex gap-4">
                                    <button type="button" class="group relative flex items-stretch justify-center text-center font-medium bg-lightprimary text-primary hover:bg-primary dark:hover:bg-primary hover:text-white rounded-lg btn-circle p-0">
                                        <span class="flex items-center gap-2 transition-all duration-150 justify-center rounded-md px-4 py-2 text-sm">
                                            <span></span>
                                        </span>
                                    </button>
                                </div>
                            </div>
                            <div data-simplebar="init">
                                <div class="simplebar-wrapper">
                                    <div class="simplebar-height-auto-observer-wrapper">
                                        <div class="simplebar-height-auto-observer"></div>
                                    </div>
                                    <div class="">
                                        <div class="">
                                            <div class="simplebar-content-wrapper" aria-label="scrollable content">
                                                <div class="simplebar-content">
                                                    <div class="border rounded-md border-border dark:border-darkborder overflow-x-auto">
                                                        <div data-testid="table-element" class="relative">
                                                            <div class="absolute left-0 top-0 -z-10 h-full w-full bg-transparent drop-shadow-md"></div>
                                                            <table class="w-full text-left text-sm text-gray-500 dark:text-gray-400">
                                                                <thead class="group/head text-sm font-medium capitalize text-dark dark:text-white border-b border-ld">
                                                                    <tr>
                                                                        <th class="px-4 dark:bg-dark dark:text-darklink text-base font-semibold py-3">
                                                                            <input type="checkbox" class="rounded border border-border dark:border-darkborder cursor-pointer text-primary checkbox">
                                                                        </th>
                                                                        <th class="px-4 dark:bg-dark dark:text-darklink text-base font-semibold py-3">ID</th>
                                                                        <th class="px-4 dark:bg-dark dark:text-darklink text-base font-semibold py-3">Title</th>
                                                                        <th class="px-4 dark:bg-dark dark:text-darklink text-base font-semibold py-3">Description</th>
                                                                        <th class="px-4 dark:bg-dark dark:text-darklink text-base font-semibold py-3">Action</th>
                                                                    </tr>
                                                                </thead>
                                                                <tbody class="group/body divide-y divide-border dark:divide-darkborder">
                                                                    <tr v-for="toDo in todos" data-testid="table-row-element" class="group/row bg-transparent">
                                                                        <td class="px-4 py-4 dark:bg-transparent whitespace-nowrap">
                                                                            <input class="rounded border border-border dark:border-darkborder cursor-pointer text-primary checkbox" type="checkbox">
                                                                        </td>
                                                                        <td class="px-4 py-4 dark:bg-transparent whitespace-nowrap lg:min-w-auto min-w-[250px]">
                                                                            <div class="flex  gap-3 items-center">
                                                                                <div class="text-no-wrap">
                                                                                    <h6 class="text-base">{{ toDo.id }}</h6>
                                                                                </div>
                                                                            </div>
                                                                        </td>
                                                                        <td class="px-4 py-4 dark:bg-transparent whitespace-nowrap">
                                                                            <p class="text-sm text-darklink font-medium">{{ toDo.title }}</p>
                                                                        </td>
                                                                        <td class="px-4 py-4 dark:bg-transparent whitespace-nowrap">
                                                                            <div class="flex gap-2 text-sm items-center text-darklink font-medium">
                                                                                <span class="flex items-center font-medium bg-success text-white text-xs h-2 w-2 p-0 rounded-md" data-testid="flowbite-badge"></span>
                                                                                {{ toDo.description }}
                                                                            </div>
                                                                        </td>
                                                                        <td class="px-4 py-4 dark:bg-transparent whitespace-nowrap">
                                                                            <div class="flex justify-center gap-3">
                                                                                <div class="w-auto" data-testid="flowbite-tooltip-target">
                                                                                    <button type="button" class="group relative flex items-stretch justify-center text-center font-medium rounded-lg btn-circle p-0 mb-2 bg-lightsuccess text-success hover:bg-success hover:text-white">
                                                                                        <span class="flex items-center gap-2 transition-all duration-150 justify-center rounded-md px-4 py-2 text-sm">
                                                                                            <nuxt-link :to="`/todo/${toDo.id}/edit`">
                                                                                            <!-- <a :href="`/todo/${toDo.ID}/edit`"> -->
                                                                                                <svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-hidden="true" role="img" class="iconify iconify--tabler" width="18" height="18" viewBox="0 0 24 24">
                                                                                                    <path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 20h4L18.5 9.5a2.828 2.828 0 1 0-4-4L4 16zm9.5-13.5l4 4"></path>
                                                                                                </svg>
                                                                                            <!-- </a> -->
                                                                                            </nuxt-link>
                                                                                        </span>
                                                                                    </button>
                                                                                </div>
                                                                                <div class="w-auto" data-testid="flowbite-tooltip-target">
                                                                                    <button @click="confirmDelete(toDo.id)" type="button" class="group relative flex items-stretch justify-center text-center font-medium bg-lighterror dark:bg-darkerror text-error hover:bg-error hover:text-white rounded-lg btn-circle p-0 mb-2">
                                                                                        <span class="flex items-center gap-2 transition-all duration-150 justify-center rounded-md px-4 py-2 text-sm">
                                                                                            <svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-hidden="true" role="img" class="iconify iconify--tabler" width="18" height="18" viewBox="0 0 24 24">
                                                                                                <path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 7h16m-10 4v6m4-6v6M5 7l1 12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2l1-12M9 7V4a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v3"></path>
                                                                                            </svg>
                                                                                        </span>
                                                                                    </button>
                                                                                </div>
                                                                            </div>
                                                                        </td>
                                                                    </tr>
                                                                </tbody>
                                                            </table>
                                                        </div>
                                                    </div>
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
    import Sidebar from '@/components/Sidebar.vue';
    import { ref, onMounted } from 'vue';
    import { useRouter } from 'vue-router';
    import axios from 'axios';
    
    const todos = ref([]);
    let socket: WebSocket;

    // Hàm fetch dữ liệu từ API
    const toDo = async () => {
        try {
            const response = await axios.get('http://localhost:8088/api/to-do-list'); // Thay URL bằng endpoint API thực tế

            console.log(response)
            todos.value = response.data; // Gán dữ liệu API trả về cho todos
        } catch (error) {
            console.error('Error fetching to-do list:', error);
        }
    };

    const deleteToDo = async (id) => {
        try {
            await axios.delete(`http://localhost:8088/api/to-do-list/${id}`);
            // Sau khi xóa, fetch lại danh sách to-do
            // Xóa to-do khỏi danh sách ngay lập tức
            // todos.value = todos.value.filter(todo => todo.id !== id);
            await toDo();
        } catch (error) {
            console.error('Error deleting to-do:', error);
        }
    };

    const confirmDelete = (id) => {
        const isConfirmed = confirm("Are you sure delete?")

        if (isConfirmed) {
            deleteToDo(id)
        }
    }

    // Tự động gọi hàm khi component được mount
    toDo();

    onMounted(() => {
        socket = new WebSocket('ws://localhost:8088/ws');

        socket.onmessage = (event) => {
            const message = JSON.parse(event.data);
            console.log(message)
            if (message.action === 'new_todo') {
                todos.value.push(message.data); // Thêm To-Do mới vào danh sách
            }

            const updatedToDo = JSON.parse(event.data);
            const index = todos.value.findIndex((todo) => todo.id === updatedToDo.id);

            if (index !== -1) {
                todos.value[index] = updatedToDo; // Cập nhật nếu tồn tại
            }
        };

        socket.onclose = () => {
            console.log('WebSocket connection closed');
        };
    });
    console.log(todos.value)
</script>
