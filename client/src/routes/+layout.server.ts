import { SERVER_URL } from "$env/static/private";
import api from "$lib/server/api";
import { z } from "zod";
import type { LayoutServerLoad } from "./$types";
import { error } from '@sveltejs/kit';

export type Favorite = {
    id: string;
    anime_id: string;
    anime_title: string;
    anime_image_url: string;
    created: string;
    updated: string;
};

const favoriteSchema = z.object({
    id: z.string(),
    anime_id: z.string(),
    anime_title: z.string(),
    anime_image_url: z.string(),
    created: z.string().refine(date => !isNaN(Date.parse(date)), {
        message: "Invalid date format"
    }),
    updated: z.string().refine(date => !isNaN(Date.parse(date)), {
        message: "Invalid date format"
    }),
});


export const load: LayoutServerLoad = (async () => {

    try {

        const response = await api<Favorite[]>(`${SERVER_URL}/favorites`);
        if (!response.success) {
            throw error(500, response.error);
        }
        const favoritesData: Favorite[] = response.data;

        favoritesData.forEach(fav => {
            try {
                favoriteSchema.parse(fav);
            } catch (e) {
                console.error("Favorite validation error:", e);
                throw error(500, "Favorite validation failed.");
            }
        });

        // convert the array of favorites to a Map
        const favorites = new Map<string, { title: string; image: string }>(
            favoritesData.map(fav => [fav.id, {
                title: fav.anime_title,
                image: fav.anime_image_url
            }])
        );

        return {
            favorites: favorites
        };

    } catch (error) {
        console.error("An error occurred while fetching favorites:", error);
        return {
            favorites: new Map() // return an empty Map in case of error
        };
    }
}) satisfies LayoutServerLoad;
