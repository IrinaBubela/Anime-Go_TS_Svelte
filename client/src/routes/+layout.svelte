<script lang="ts">
    import "../app.css";
    import type { LayoutData } from "./$types";
    import Anime from "./anime.svelte";

    export let data: LayoutData;
</script>

<div class="grid grid-cols-2 p-4">
    <div>
        <slot />
    </div>
    <div>
        {#each [...data.favorites] as [key, value]}
            <div class="my-2">
                <Anime title={value.title} mal_id={key} image={value.image} />
            </div>
        {/each}
    </div>
</div>
