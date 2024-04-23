"use client";

import { useEffect, useState } from "react";
import { useToast } from "@/components/ui/use-toast";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Send } from "lucide-react";
import Navbar from "@/components/navbar";

interface message {
  sender: string;
  content: string;
}

export default function Chat() {
  const [messages, setMessages] = useState<message[]>([]);
  const [ws, setWs] = useState<WebSocket | null>(null);
  const { toast } = useToast();
  const [inputValue, setInputValue] = useState("");
  const [username, setUserName] = useState("");

  useEffect(() => {
    const ws = new WebSocket("ws://localhost:8080/ws");

    ws.onopen = () => {
      console.log("connected");
      toast({
        title: "Connection established!",
        description: "You can now send messages!",
      });
      let name = localStorage.getItem("username")!;
      setUserName(name);
    };

    ws.onmessage = (msg) => {
      console.log("message received: ", msg.data);
      const newMessage = JSON.parse(msg.data);
      setMessages((prevMessages) => [...prevMessages, newMessage]);
    };

    ws.onclose = () => {
      console.log("disconnected");
    };

    setWs(ws);

    return () => {
      ws.close();
    };

  }, []);

  const renderMessages = () => {
    return messages.map((message: message, index: number) => (
      <div
        key={index}
        className={`max-w-[70%] space-y-1.5 ${
          message.sender === username ? "justify-end" : ""
        }`}
      >
        <div
          className={`rounded-lg ${
            message.sender === username
              ? "bg-white px-4 py-3 shadow-sm dark:bg-gray-800 dark:text-gray-200"
              : "bg-gray-900 px-4 py-3 text-white"
          }`}
        >
          <p>{message.content}</p>
        </div>
      </div>
    ));
  };

  const sendMessage = () => {
    if (ws && inputValue) {
      const message = {
        sender: username,
        content: inputValue,
      };
      ws.send(JSON.stringify(message));
      toast({
        title: "Message sent!",
      });
    }
  };

  return (
    <>
      <Navbar />
      <div className="flex h-[calc(100vh-75px)] w-full flex-col bg-gray-100 dark:bg-black">
        <div className="flex-1 overflow-y-auto p-6">
          <div className="grid gap-4">{renderMessages()}</div>
        </div>
        <div className="flex h-16 items-center justify-between border-t bg-white px-6 shadow-sm dark:border-gray-800 dark:bg-gray-950">
          <Input
            className="flex-1 bg-transparent pr-4"
            placeholder="Type your message..."
            type="text"
            value={inputValue}
            onChange={(e) => setInputValue(e.target.value)}
          />
          <Button size="icon" variant="ghost" onClick={sendMessage}>
            <Send className="h-5 w-5" />
          </Button>
        </div>
      </div>
    </>
  );
}
