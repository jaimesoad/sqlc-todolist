<script lang="ts">
    import { NewTodo } from "./requests";
    import type { todo } from "./Types"

    export let datos: Promise<todo[]>

    let value: todo = {
        id: 0,
        content: "",
        done: false
    }
</script>

<form>
    <input type="text" placeholder="New Todo" bind:value={value.content} />
    <button on:click={() => datos = NewTodo(value)}>Save</button>
</form>

<style>
    form {
        display: flex;
        width: 30.625rem;
        height: 4rem;
        justify-content: center;
        align-items: center;
        gap: 1rem;
    }

    input {
        width: 23rem;
        height: 44px;
        flex-shrink: 0;
        border-radius: 1rem;
        border: 2px solid #ababab;
        background: #dfdfdf;
        outline: none;
        padding: 0 1rem;
        font-size: 1rem;
        font-family: Arial, Helvetica, sans-serif;
    }

    input::placeholder {
        color: #8A8A8A;
    }

    button {
        width: 4.625rem;
        height: 3rem;
        flex-shrink: 0;
        border-radius: 1rem;
        border: 2px solid #1865ad;
        background: #389cf9;
        color: white;
        font-size: 1rem;
        font-family: Arial, Helvetica, sans-serif;
    }
</style>
