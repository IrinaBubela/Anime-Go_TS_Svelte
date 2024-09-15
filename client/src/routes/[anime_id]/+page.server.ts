import api from "$lib/server/api";
import { error } from "@sveltejs/kit";
import type { Actions, PageServerLoad } from "./$types";
import { SERVER_URL } from "$env/static/private";
import { toast } from 'svelte-sonner';

export type Anime = {
    data: {
        mal_id: number;
        title: string;
        url: string;
        synopsis: string;
        images: {
            webp: {
                image_url: string;
                small_image_url: string;
                large_image_url: string;
            };
        };
    };
};

export const load = (async ({ params }) => {
    const id = params.anime_id;
    const anime = await api<Anime>(
        `https://api.jikan.moe/v4/anime/${id}`,
    );
    if (!anime.success) {
        console.error("Failed to fetch anime", anime.error);
        throw error(500, "Failed to fetch anime");
    }
    console.debug("anime", anime.data);
    return {
        anime: anime.data.data,
    };
}) satisfies PageServerLoad;

export const actions = {
    addToFavorites: async ({ request }) => {
        const form = await request.formData();

        // validate form
        const mal_id = form.get("mal_id") as unknown as number;
        const title = form.get("title") as unknown as string;
        const image = form.get("image") as unknown as string;

        try {
            const response = await api(`${SERVER_URL}/favorites`,
                {
                    method: 'POST',
                    body: {
                        id: mal_id.toString(),
                        anime_id: mal_id.toString(),
                        anime_title: title,
                        anime_image_url: image,
                        created: new Date().toISOString(),
                        updated: new Date().toISOString(),
                    }
                }
            );

            if (response.success) {
                toast('Favorite added successfully!');
            } else {
                if (response.error === 'Maximum limit of 5 favorites reached') {
                    toast('You have reached the limit of 5 favorites.');
                } else {
                    toast('An error occurred while adding the favorite.');
                }
            }
        } catch (error) {
            toast('An unexpected error occurred.');
        }
    },
    deleteFavorite: async ({ request }) => {
        const form = await request.formData();
        const id = form.get("mal_id") as unknown as string;

        try {
            const response = await api(`${SERVER_URL}/favorites/${id}`, {
                method: 'DELETE',
            });

            if (response.success) {
                toast('Favorite deleted successfully!');
            } else {
                toast('An error occurred while deleting the favorite.');
            }
        } catch (error) {
            toast('An unexpected error occurred while deleting the favorite:', error);
        }
    }
} satisfies Actions;
