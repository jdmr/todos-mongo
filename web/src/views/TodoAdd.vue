<template>
    <div
        class="container mx-auto max-w-prose p-2 flex flex-col gap-4 items-center justify-center"
    >
        <h1 class="text-5xl tracking-widest">TODOS</h1>
        <form
            action="addTodo"
            class="grid gap-4 w-full"
            method="post"
            @submit.prevent="addTodo"
        >
            <div class="grid gap-1">
                <label for="name">Name</label>
                <input
                    id="name"
                    v-model="todo.name"
                    type="text"
                    name="name"
                    class="border border-gray-300 rounded p-2 bg-gray-800 tracking-wider"
                />
            </div>
            <div class="grid gap-1">
                <label for="description">Description</label>
                <textarea
                    id="description"
                    v-model="todo.description"
                    name="description"
                    class="border border-gray-300 rounded p-2 bg-gray-800 tracking-wider"
                ></textarea>
            </div>
            <div class="grid mt-2">
                <button
                    type="submit"
                    class="bg-blue-500 text-white p-4 rounded text-2xl"
                >
                    Add
                </button>
            </div>
        </form>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { Todo } from '../types/todos';

const todo = ref({} as Todo);
const adding = ref(false);
const router = useRouter();

const addTodo = () => {
    if (adding.value) {
        return;
    }

    if (!todo.value.name) {
        alert('Name is required');
        return;
    }

    try {
        adding.value = true;
        // post request to add todo
        fetch('/api/v1/todos', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(todo.value)
        });

        todo.value = {} as Todo;
        router.go(-1);
    } catch (err) {
        console.error('Error adding todo', err);
        alert('Error adding todo');
    } finally {
        adding.value = false;
    }
};
</script>
