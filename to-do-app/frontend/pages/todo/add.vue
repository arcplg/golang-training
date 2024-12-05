
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
                                    <h4 class="font-semibold text-xl text-dark dark:text-white mb-3">Add To do list</h4>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div data-testid="flowbite-card" class="flex rounded-md bg-white dark:bg-dark relative w-full break-words flex-col card p-6 dark:shadow-dark-md shadow-md undefined" style="border-radius:7px">
                        <div class="flex h-full flex-col justify-start gap-0 p-0">
                            <div data-simplebar="init">
                                <div class="simplebar-wrapper">
                                    <div class="grid grid-cols-12 gap-6">
                                        <div class="col-span-12">
                                            <div data-testid="flowbite-card" class="flex rounded-md bg-white dark:bg-dark relative w-full break-words flex-col card undefined dark:shadow-dark-md shadow-md p-0" style="border-radius: 7px;">
                                                <div class="flex h-full flex-col justify-start gap-0 p-0">
                                                    <form @submit.prevent="addToDo">
                                                        <div class="pt-4 p-6">
                                                            <div class="grid grid-cols-6 gap-[1.875rem]">
                                                                <div class="col-span-12">
                                                                    <div class="mb-2 block">
                                                                        <label class="text-sm font-semibold text-gray-900 dark:text-white" data-testid="flowbite-label" for="Default Text">Title</label>
                                                                    </div>
                                                                    <div class="flex form-control">
                                                                        <div class="relative w-full">
                                                                            <input v-model="toDo.title" class="block w-full border disabled:cursor-not-allowed disabled:opacity-50 border-gray-300 bg-gray-50 text-gray-900 focus:border-cyan-500 focus:ring-cyan-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-cyan-500 dark:focus:ring-cyan-500 p-2.5 text-sm rounded-lg" id="default" placeholder="Marcal" type="text">
                                                                        </div>
                                                                    </div>
                                                                </div>
                                                                <div class="col-span-12">
                                                                    <div class="mb-2 block">
                                                                        <label class="text-sm font-semibold text-gray-900 dark:text-white" data-testid="flowbite-label" for="comment">Description</label>
                                                                    </div>
                                                                    <textarea v-model="toDo.description" class="block w-full rounded-md text-sm disabled:cursor-not-allowed disabled:opacity-50 bg-transparent border border-ld text-dark focus:border-primary focus:ring-0 dark:text-white dark:placeholder-gray-400 dark:focus:border-primary dark:focus:ring-0 form-control-textarea" id="comment" placeholder="Leave a comment..." required="" rows="4"></textarea>
                                                                </div>
                                                                <div class="col-span-12">
                                                                    <button type="submit" class="group relative flex items-stretch justify-center p-0.5 text-center font-medium bg-primary text-white rounded-lg">
                                                                        <span class="flex items-center gap-2 transition-all duration-150 justify-center rounded-md px-4 py-2 text-sm">Submit</span>
                                                                    </button>
                                                                </div>
                                                            </div>
                                                        </div>
                                                    </form>
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

    import { ref } from 'vue';
    import { useRouter } from 'vue-router';
    import axios from 'axios'

    const toDo = ref({
        title: '',
        description: '',
        is_completed: true
    })

    const router = useRouter()

    const addToDo = async () => {
        try {
            const response = await axios.post('http://localhost:8088/api/to-do-list', toDo.value); // Use the correct Axios payload format
            console.log('To do add:', response.data);

            alert('Add successfully!');

            // Navigate to the home page after registration
            router.push('/homepage');
        } catch (error) {
            console.error('Error add to do:', error.response?.data || error.message);
        }
    };
</script>
