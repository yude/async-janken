import { NextResponse } from 'next/server';
import { KVNamespace } from "@cloudflare/workers-types";

export const runtime = 'edge';

const { MY_KV_STORE } = process.env as unknown as {
    MY_KV_STORE: KVNamespace;
};

export async function GET(request: Request) {
  
  return NextResponse.json(id);
}
