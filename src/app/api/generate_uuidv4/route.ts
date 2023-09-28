import { NextResponse } from 'next/server';
import { v4 } from "uuid";

export const runtime = 'edge';

export async function GET(request: Request) {
    return NextResponse.json({result: v4()})
}
