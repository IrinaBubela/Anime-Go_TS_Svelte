<script lang="ts">
    import { enhance } from "$app/forms";
    import { Toaster } from "svelte-sonner";
    import Anime from "../anime.svelte";
    import type { PageData } from "./$types";
    import { onMount } from "svelte";
    import { gsap } from "gsap";

    export let data: PageData;

    onMount(() => {
        gsap.from(".anime-detail", {
            opacity: 0,
            y: 20,
            duration: 0.5,
            ease: "power2.out"
        });
    });
</script>

<Toaster />

<div class="p-4 anime-detail">
    <Anime
        title={data.anime.title}
        mal_id={data.anime.mal_id}
        image={data.anime.images.webp.image_url}
    />
    <p class="mt-4 text-gray-700">{data.anime.synopsis}</p>
    <form action="?/addToFavorites" method="post" use:enhance>
        <input type="hidden" name="mal_id" value={data.anime.mal_id} />
        <input type="hidden" name="title" value={data.anime.title} />
        <input type="hidden" name="image" value={data.anime.images.webp.image_url} />
        <button class="mt-4 rounded bg-gray-300 p-4 hover:bg-gray-400 transition-colors duration-300" type="submit">
            Add to favorites
        </button>
    </form>
    <div class="mt-6">
        <a class="rounded bg-gray-300 p-4 inline-block hover:bg-gray-400 transition-colors duration-300" href="/">Go back to list</a>
    </div>
</div>
