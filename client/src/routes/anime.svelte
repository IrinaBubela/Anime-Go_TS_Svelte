<script lang="ts">
    export let title: string;
    export let mal_id: number;
    export let image: string;
</script>

<h1>{title}</h1>

<a href="/{mal_id}">
    <img
        class="hover:cursor-pointer hover:opacity-50"
        src={image}
        alt={title}
        height="200px"
        width="200px"
    />
</a>

<form action="?/deleteFavorite" method="POST">
    <input type="hidden" name="mal_id" value={mal_id} />
    <input type="hidden" name="_method" value="DELETE">
    <button
        class="mt-4 rounded bg-red-500 px-3 py-1 text-white" type="submit">
        Remove from Favorites
    </button>
</form>