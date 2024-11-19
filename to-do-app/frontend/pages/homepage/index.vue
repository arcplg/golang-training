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
                                    <ol class="flex items-center whitespace-nowrap" aria-label="Breadcrumb">
                                        <li class="flex items-center">
                                            <a class="opacity-80 text-sm text-link dark:text-darklink leading-none" href="@@webRoot/main/index.html">Home</a>
                                        </li>
                                        <li>
                                            <div class="p-0.5 rounded-full bg-dark dark:bg-darklink mx-2.5 flex items-center"></div>
                                        </li>
                                        <li class="flex items-center text-sm text-link dark:text-darklink leading-none" aria-current="page">To do list</li>
                                    </ol>
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
                            <div data-simplebar="init" class="max-h-[580px]">
                                <div class="simplebar-wrapper">
                                    <div class="simplebar-height-auto-observer-wrapper">
                                        <div class="simplebar-height-auto-observer"></div>
                                    </div>
                                    <div class="simplebar-mask">
                                        <div class="simplebar-offset">
                                            <div class="simplebar-content-wrapper" tabindex="0" role="region" aria-label="scrollable content">
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
                                                                        <th class="px-4 dark:bg-dark dark:text-darklink text-base font-semibold py-3">Products</th>
                                                                        <th class="px-4 dark:bg-dark dark:text-darklink text-base font-semibold py-3">Date</th>
                                                                        <th class="px-4 dark:bg-dark dark:text-darklink text-base font-semibold py-3">Status</th>
                                                                        <th class="px-4 dark:bg-dark dark:text-darklink text-base font-semibold py-3">Price</th>
                                                                        <th class="px-4 dark:bg-dark dark:text-darklink text-base font-semibold py-3">Action</th>
                                                                    </tr>
                                                                </thead>
                                                                <tbody class="group/body divide-y divide-border dark:divide-darkborder">
                                                                    <tr data-testid="table-row-element" class="group/row bg-transparent">
                                                                        <td class="px-4 py-4 dark:bg-transparent whitespace-nowrap">
                                                                            <input class="rounded border border-border dark:border-darkborder cursor-pointer text-primary checkbox" type="checkbox">
                                                                        </td>
                                                                        <td class="px-4 py-4 dark:bg-transparent whitespace-nowrap lg:min-w-auto min-w-[250px]">
                                                                            <div class="flex  gap-3 items-center">
                                                                                <img src="#" alt="icon" width="56" height="56" class="h-14 w-14 rounded-full">
                                                                                <div class="text-no-wrap">
                                                                                    <h6 class="text-base">Cute Soft Teddybear</h6>
                                                                                    <p class="text-sm text-darklink">toys</p>
                                                                                </div>
                                                                            </div>
                                                                        </td>
                                                                        <td class="px-4 py-4 dark:bg-transparent whitespace-nowrap">
                                                                            <p class="text-sm text-darklink font-medium">Mon, Nov 18 2024</p>
                                                                        </td>
                                                                        <td class="px-4 py-4 dark:bg-transparent whitespace-nowrap">
                                                                            <div class="flex gap-2 text-sm items-center text-darklink font-medium">
                                                                                <span class="flex items-center font-medium bg-success text-white text-xs h-2 w-2 p-0 rounded-md" data-testid="flowbite-badge"></span>InStock
                                                                            </div>
                                                                        </td>
                                                                        <td class="px-4 py-4 dark:bg-transparent whitespace-nowrap">
                                                                            <h5 class="text-base">$285</h5>
                                                                        </td>
                                                                        <td class="px-4 py-4 dark:bg-transparent whitespace-nowrap">
                                                                            <span class="h-9 w-9 flex justify-center items-center rounded-full hover:bg-lightprimary hover:text-primary cursor-pointer" aria-expanded="false" aria-haspopup="menu" id=":r526:">
                                                                                <svg stroke="currentColor" fill="none" stroke-width="2" viewBox="0 0 24 24" aria-hidden="true" height="22" width="22" xmlns="http://www.w3.org/2000/svg">
                                                                                    <path stroke-linecap="round" stroke-linejoin="round" d="M12 5v.01M12 12v.01M12 19v.01M12 6a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2z"></path>
                                                                                </svg>
                                                                            </span>
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
</script>
