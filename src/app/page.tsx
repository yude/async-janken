"use client";
import { useState, useEffect } from "react";
import { Match } from "@/app/types/match";

export default function Home() {
  const [hand, setHand] = useState(0);

  const [matches, setMatches] = useState<Match[]>();

  return (
    <main className="text-center max-w-lg mx-auto mt-10">
      <div className="text-2xl">非同期じゃんけん</div>
      <div className="bg-slate-600 text-slate-300 p-5">
        <p className="text-xl underline">説明</p>
        <p>あなたが、どの手を出すか決めます。そのときに...</p>
        <div className="border-2 border-dotted"></div>
        <p className="font-bold">まだ、相手がいなかったとき</p>
        <p>
          <span className="inline-block">
            再度アクセスしたときに誰かが手を決めていれば、
          </span>
          <span className="inline-block">勝敗が分かります。</span>
        </p>
        <p className="font-bold">誰かが、既に手を決めていたとき</p>
        <p>すぐにじゃんけんが行われ、勝敗が分かります。</p>
      </div>
      <div className="m-5">
        {/* 
            hand:
              - 0: ✋
              - 1: ✊
              - 2: ✌
          */}
        <button
          className="text-6xl bg-pink-400 hover:bg-pink-300 text-white rounded px-10 py-8"
          onClick={() => {
            setHand(0);
          }}
        >
          ✊
        </button>
        <button
          className="text-6xl bg-pink-400 hover:bg-pink-300 text-white rounded px-10 py-8"
          onClick={() => {
            setHand(1);
          }}
        >
          ✌
        </button>
        <button
          className="text-6xl bg-pink-400 hover:bg-pink-300 text-white rounded px-10 py-8"
          onClick={() => {
            setHand(2);
          }}
        >
          ✋
        </button>
      </div>
      <div className="bg-slate-600 text-slate-300 p-5 mt-10">
        <p className="text-xl underline">対戦履歴</p>
        <p>wip</p>
      </div>
    </main>
  );
}
