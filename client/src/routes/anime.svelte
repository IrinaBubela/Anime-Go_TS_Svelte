<script lang="ts">
    import { enhance } from "$app/forms";

    export let title: string;
    export let mal_id: number;
    export let image: string;
</script>

<div class="anime-screen max-w-xs mx-auto bg-white shadow-lg rounded-lg overflow-hidden p-4">
    <h1 class="text-xl font-semibold text-center mb-4">{title}</h1>

    <a href="/{mal_id}" class="block text-center flex justify-center">
        <img
            class="transition-transform transform hover:scale-105 hover:opacity-80 rounded-lg"
            src={image}
            alt={title}
            height="200px"
            width="200px"
        />
    </a>

    <form action="?/deleteFavorite" method="post" use:enhance class="mt-4 flex justify-center">
        <input type="hidden" name="mal_id" value={mal_id} />
        <button
            class="rounded bg-red-600 hover:bg-red-700 px-4 py-2 text-white font-semibold transition duration-300"
            type="submit">
            Remove from Favorites
        </button>
    </form>
</div>
