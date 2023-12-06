<script lang="ts">
    import type { todo } from "./Types";
    import { DeleteTodo, ToggleTodo } from "./requests";

    export let data: todo;
    $: color = data.done ? "done" : "not_done"

    let nodeRef: Element

    async function toggle() {
        data.done = await ToggleTodo(data)
    }
    async function destroy() {
        if (!await DeleteTodo(data)) return

        nodeRef.parentNode?.removeChild(nodeRef)
    }
</script>

<div id="wrapper" bind:this={nodeRef}>
    <p>{data.content}</p>
    <button id="{color}" class="btn" on:click={() => toggle()}></button>
    <button id="exit" on:click={() => destroy()}>&times</button>
</div>

<style>
    #wrapper {
        position: relative;
        width: 15rem;
        background: #dddddd;
        color: black;
        display: flex;
        align-items: center;
        justify-content: space-between;
        padding: 1.69rem 1.25rem;
        border-radius: 15px;
        margin-top: 1rem;
    }

    #wrapper p {
        margin: 0;
        padding: 0;
        color: #000;
        font-family: Arial, Helvetica, sans-serif;
        font-size: 1rem;
        font-style: normal;
        font-weight: 400;
        line-height: normal;
    }

    .btn {
        height: 1.875rem;
        aspect-ratio: 1/1;
        border-radius: 5px;
        border: none;
        text-align: center;
        font-size: 1rem;
        transition: 200ms;
    }

    .btn:active {
        transform: rotate(90deg);
    }

    #done {
        background: #48bec6;
    }

    #not_done {
        background: #e74343;
    }

    #exit {
        padding: 0;
        height: 1rem;
        aspect-ratio: 1/1;
        border-radius: 50%;
        border: none;
        background: #e74343;
        position: absolute;
        top: -0.62rem;
        right: -0.62rem;
        overflow: hidden;
        color: #00000000;
        transition: 200ms;
        display: flex;
        align-items: center;
        justify-content: center;
        font-size: 1.5rem;
    }

    #exit:hover {
        color: #7c2424;
    }
</style>
