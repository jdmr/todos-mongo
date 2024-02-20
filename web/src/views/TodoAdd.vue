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
            <div class="grid gap-1">
                <label for="dueDate">Due Date</label>
                <input
                    id="dueDate"
                    v-model="todo.dueDate"
                    type="date"
                    name="name"
                    class="border border-gray-300 rounded p-2 bg-gray-800 tracking-wider"
                />
            </div>
            <div class="grid gap-1">
                <label for="status">Status</label>
                <select
                    id="status"
                    v-model="status"
                    name="status"
                    class="border border-gray-300 rounded p-2 bg-gray-800 tracking-wider"
                >
                    <option value="1">Pending</option>
                    <option value="2">In Progress</option>
                    <option value="3">Completed</option>
                </select>
            </div>
            <div class="grid gap-1">
                <label for="priority">Priority</label>
                <select
                    id="priority"
                    v-model="priority"
                    name="priority"
                    class="border border-gray-300 rounded p-2 bg-gray-800 tracking-wider"
                >
                    <option value="1">Low</option>
                    <option value="2">Medium</option>
                    <option value="3">High</option>
                </select>
            </div>
            <div class="grid gap-1">
                <label for="user">User</label>
                <select
                    id="user"
                    v-model="user"
                    name="user"
                    class="border border-gray-300 rounded p-2 bg-gray-800 tracking-wider"
                >
                    <option value="1">John</option>
                    <option value="2">Sarah</option>
                    <option value="3">Will</option>
                    <option value="4">Mary</option>
                </select>
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
import { Todo, User } from '../types/todos';

const todo = ref({} as Todo);
const adding = ref(false);
const router = useRouter();
const status = ref('1');
const priority = ref('1');
const user = ref('1');
const statuses = ref([
    { id: '1', name: 'Pending' },
    { id: '2', name: 'In Progress' },
    { id: '3', name: 'Completed' }
]);
const priorities = ref([
    { id: '1', name: 'Low' },
    { id: '2', name: 'Medium' },
    { id: '3', name: 'High' }
]);
const users = ref([
    { id: '1', name: 'John' },
    { id: '2', name: 'Sarah' },
    { id: '3', name: 'Will' },
    { id: '4', name: 'Mary' }
] as User[]);

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
        todo.value.status = statuses.value.find((s) => s.id === status.value);
        todo.value.priority = priorities.value.find(
            (p) => p.id === priority.value
        );
        todo.value.user = users.value.find((u) => u.id === user.value);

        // post request to add todo
        fetch('/api/v1/todos', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(todo.value)
        });

        todo.value = {} as Todo;
        router.push('/');
    } catch (err) {
        console.error('Error adding todo', err);
        alert('Error adding todo');
    } finally {
        adding.value = false;
    }
};
</script>
