<script lang="ts">
    import { enhance } from "$app/forms";
    import Anime from "../anime.svelte";
    import type { PageData } from "./$types";
    import { onMount } from "svelte";
    import { gsap } from "gsap";
    import { Toaster, toast } from "svelte-sonner";

    export let data: PageData;
    export let form: { error: string; success: string } = {
        error: "",
        success: "",
    };

    $: if (form) {
        if (form.success) {
            toast.success(form.success);
        } else if (form.error) {
            toast.error(form.error);
        }
    }

    onMount(() => {
        gsap.from(".anime-detail", {
            opacity: 0,
            y: 20,
            duration: 0.5,
            ease: "power2.out",
        });
    });
</script>

<Toaster position="top-center" expand={true} richColors />

<div class="anime-detail p-4">
    <Anime
        title={data.anime.title}
        mal_id={data.anime.mal_id}
        image={data.anime.images.webp.image_url}
    />
    <p class="mt-4 text-gray-700">{data.anime.synopsis}</p>
    <form action="?/addToFavorites" method="post" use:enhance>
        <input type="hidden" name="mal_id" value={data.anime.mal_id} />
        <input type="hidden" name="title" value={data.anime.title} />
        <input
            type="hidden"
            name="image"
            value={data.anime.images.webp.image_url}
        />
        <button
            class="mt-4 rounded bg-gray-300 p-4 transition-colors duration-300 hover:bg-gray-400"
            type="submit"
        >
            Add to favorites
        </button>
    </form>
    <div class="mt-6">
        <a
            class="inline-block rounded bg-gray-300 p-4 transition-colors duration-300 hover:bg-gray-400"
            href="/"
        >
            Go back to list
        </a>
    </div>
</div>
