<script lang="ts">
    import { LastTenTodos } from "./requests";
    import Todo from "./todo.svelte";
    import Form from "./form.svelte";

    let datos = LastTenTodos();
</script>

<nav>
    <Form bind:datos />
</nav>

<div id="todos">
    {#await datos then todos}
        {#each todos as data}
            <Todo bind:data />
        {/each}
    {/await}
</div>

<style>
    :global(body) {
        background: #1e1e1e;
        color: #dddddd;
        padding: 0;
        margin: 0;
    }

    #todos {
        width: fit-content;
        margin: 0 auto;
        margin-top: 5rem;
    }

    nav {
        width: 100%;
        display: flex;
        justify-content: center;
        background: #3939396c;
        backdrop-filter: blur(15px);
        position: fixed;
        top: 0;
        z-index: 2;
    }
</style>
