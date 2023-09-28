import { NextResponse } from 'next/server';
import { collection, DocumentData, getDocs } from "firebase/firestore";
import { query, orderBy, limit } from "firebase/firestore";

import { db } from "@/app/lib/firestore";
import { Match } from "@/app/types/match"

export const runtime = 'edge';

export async function GET(request: Request) {
    const querySnapshot = await getDocs(
        query(
            collection(db, "matches"),
            orderBy('datetime', 'desc'),
            limit(10)
        )
    );

    const matches: Match[] = [];

    querySnapshot.forEach((q) => {
        const match = {
            players: {
                0: q.data().player_0,
                1: q.data().player_1,
            },
            winner: q.data().winner,
            timestamp: q.data().timestamp,
        }
        matches.push(match)
    })

    return NextResponse.json(matches);
}
