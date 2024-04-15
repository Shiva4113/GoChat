"use client"

import { Button } from "@/components/ui/button";
import { connect, sendMsg } from "@/lib/api";
import { useEffect } from "react";

export default function Home() {
  useEffect(() => {
    connect();
  }, []);

  console.log(process.env.BACKEND);

  function handleSendMsg() {
    console.log("Message sent");
    sendMsg("Hello World");
  }

  return (
    <main className="text-center my-auto">
      <Button onClick={handleSendMsg}>Send</Button>
    </main>
  );
}
