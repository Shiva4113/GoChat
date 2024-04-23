"use client";

import { useEffect } from "react";
import { useRouter } from "next/navigation";
import { Lamp } from "@/components/lamp";

function App() {
  const router = useRouter();

  useEffect(() => {
    const timer = setTimeout(() => {
      router.push("/login");
    }, 2000);

    return () => clearTimeout(timer);
  }, [router]);

  return <Lamp />;
}

export default App;
