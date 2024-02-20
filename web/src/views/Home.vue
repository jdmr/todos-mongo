<template>
    <div class="p-2 flex flex-col gap-4">
        <h1 class="text-5xl tracking-widest">TODOS</h1>
        <div class="flex">
            <RouterLink
                to="/add-todo"
                class="py-2 px-4 bg-blue-500 text-white rounded flex items-center"
                ><i class="i-mdi:plus h-6 w-6"></i> Todo</RouterLink
            >
        </div>
    </div>
</template>
