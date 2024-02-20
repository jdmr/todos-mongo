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
        <div class="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
            <div
                v-for="todo in todos"
                :key="todo.id"
                class="border rounded shadow p-2 flex flex-col gap-2 justify-between"
            >
                <div>
                    <div>
                        <div>Name</div>
                        <div>{{ todo.name }}</div>
                    </div>
                    <div>
                        <div>Description</div>
                        <div>{{ todo.description }}</div>
                    </div>
                    <div>
                        <div>Status</div>
                        <div>{{ todo.status?.name }}</div>
                    </div>
                    <div>
                        <div>Priority</div>
                        <div>{{ todo.priority?.name }}</div>
                    </div>
                    <div>
                        <div>User</div>
                        <div>{{ todo.user?.name }}</div>
                    </div>
                    <div>
                        <div>Due</div>
                        <div>{{ todo.dueDate }}</div>
                    </div>
                </div>
                <div>
                    <button
                        class="bg-red-500 text-white p-2 rounded"
                        @click="deleteTodo($event, todo.id)"
                    >
                        Delete
                    </button>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { Todo } from '../types/todos';

const todos = ref([] as Todo[]);

const getTodos = async () => {
    const response = await fetch('/api/v1/todos');
    todos.value = await response.json();
    console.log(todos.value);
};

const deleteTodo = async (evt: Event, id: string) => {
    evt.preventDefault();
    let el = evt.target as HTMLButtonElement;
    try {
        if (el) {
            el.disabled = true;
        }
        await fetch(`/api/v1/todos/${id}`, { method: 'DELETE' });
        todos.value = todos.value.filter((todo) => todo.id !== id);
    } catch (error) {
        console.error(error);
        if (el) {
            el.disabled = false;
        }
    }
};

onMounted(getTodos);
</script>
