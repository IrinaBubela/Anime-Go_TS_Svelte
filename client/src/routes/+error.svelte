<script>
    import "../app.css";
    export let error;
</script>

<div
    class="flex flex-col items-center justify-center mt-6"
>
    <h1 class="my-4 text-xl font-bold">
        {error || "An unexpected error occurred."}
    </h1>

    <button
        on:click={() => (window.location.href = "/")}
        class="mt-4 rounded-md bg-red-500 px-6 py-3 font-semibold text-white shadow-lg transition duration-300 hover:bg-red-600"
    >
        Go back to home
    </button>
</div>
