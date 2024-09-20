<script lang="ts">
    import { onMount } from "svelte";
    import type { PageData } from "./$types";
    import Anime from "./anime.svelte";
    import { gsap, ScrollTrigger } from "$lib/gsap";

    export let data: PageData;
    let loading = true;

    onMount(() => {
        ScrollTrigger.refresh();
        for (let i = 0; i < data.recommended.data.length; i++) {
            gsap.from(`#anime-${i}`, {
                x: i % 2 === 0 ? 200 : -200,
                opacity: 0,
                duration: 3,
                ease: "power4.out",
                scrollTrigger: {
                    trigger: `#anime-${i}`,
                    start: "top 80%",
                },
            });
        }
        loading = false;
    });
</script>

{#if loading}
    <div
        class="spinner absolute left-0 top-0 flex h-full w-full items-center justify-center bg-gray-900"
    >
        <svg
            class="mr-3 h-10 w-10 animate-spin text-white"
            xmlns="http://www.w3.org/2000/svg"
            fill="white"
            viewBox="0 0 24 24"
            stroke="currentColor"
        >
            <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M12 6v6m0 0v6m0-6h6m-6 0H6"
            />
        </svg>
    </div>
{/if}

{#each data.recommended.data as recommendation, i}
    <div
        class="mb-4 grid grid-cols-2 rounded border border-gray-800"
        id={`anime-${i}`}
    >
        <h1 class="col-span-2 border-b border-gray-800 p-2">
            {recommendation.content}
        </h1>
        {#each recommendation.entry as subRecommendation}
            <div class="mb-2 p-2">
                <Anime
                    title={subRecommendation.title}
                    mal_id={subRecommendation.mal_id}
                    image={subRecommendation.images.webp.image_url}
                />
            </div>
        {/each}
    </div>
{/each}

<style>
    .spinner {
        z-index: 999;
    }
</style>
