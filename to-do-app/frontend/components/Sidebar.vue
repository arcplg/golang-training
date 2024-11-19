<template>
    <div class="xl:block hidden">
        <div class="flex ">
            <nav aria-label="Sidebar with multi-level dropdown example" class="h-full w-64 fixed menu-sidebar bg-white dark:bg-dark z-[6] border-r rtl:border-l border-border dark:border-darkborder">
                <div class="bg-white dark:bg-dark rounded-none w-[270px] border-r-4 border-primary dark:border-darkborder h-full">
                    <div class="px-6 flex items-center brand-logo overflow-hidden">
                        <a href="/">
                            <img alt="logo" loading="lazy" width="174" height="26" decoding="async" data-nimg="1" class="block dark:hidden rtl:scale-x-[-1]" style="color:transparent" src="/assets/images/dark-logo.31875ba3.svg">
                            <img alt="logo" loading="lazy" width="174" height="26" decoding="async" data-nimg="1" class="hidden dark:block rtl:scale-x-[-1]" style="color:transparent" src="/assets/images/light-logo.48dc618e.svg">
                        </a>
                    </div>
                    <div data-simplebar="init" class="h-[calc(100vh_-_170px)]">
                        <div class="simplebar-wrapper">
                            <div class="simplebar-height-auto-observer-wrapper">
                                <div class="simplebar-height-auto-observer"></div>
                            </div>
                            <div class="simplebar-mask">
                                <div class="simplebar-offset">
                                    <div class="simplebar-content-wrapper" tabindex="0" role="region" aria-label="scrollable content">
                                        <div class="simplebar-content">
                                            <div class="px-6" data-testid="flowbite-sidebar-items">
                                                <ul data-testid="flowbite-sidebar-item-group" class="mt-4 space-y-0 border-t border-ld pt-4 first:mt-0 first:border-t-0 first:pt-0 sidebar-nav sidebar-nav">
                                                <h5 class="text-link font-bold text-xs dark:text-darklink caption">
                                                    <span class="hide-menu leading-21">HOME</span>
                                                    <span></span>
                                                </h5>
                                                <li class="">
                                                    <a aria-labelledby="flowbite-sidebar-item-:R2gpivaja:" class="flex items-center justify-center rounded-md px-2.5 py-2.5 mb-0.5 gap-3 !text-sm text-start leading-[normal] font-normal hover:bg-lightprimary hover:text-primary dark:text-darklink dark:hover:text-primary text-link bg-transparent group/link" href="#">
                                                        <span data-testid="flowbite-sidebar-item-content" id="flowbite-sidebar-item-:R2gpivaja:" class="flex-1 whitespace-nowrap px-0 leading-21">
                                                            <span class="flex gap-3.5 align-center items-center leading-normal truncate">
                                                                <span></span>
                                                                <div class="max-w-36 truncate hide-menu flex-1">To do</div>
                                                            </span>
                                                        </span>
                                                    </a>
                                                </li>
                                            </ul>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                            <div class="simplebar-placeholder"></div>
                        </div>
                        <div class="simplebar-track simplebar-horizontal">
                            <div class="simplebar-scrollbar"></div>
                        </div>
                        <div class="simplebar-track simplebar-vertical">
                            <div class="simplebar-scrollbar"></div>
                        </div>
                    </div>
                </div>
            </nav>
        </div>
    </div>
</template>
