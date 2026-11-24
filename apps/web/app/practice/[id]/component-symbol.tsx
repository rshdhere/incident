import Image from "next/image";
import type { Component } from "./diagram-model";

export default function ComponentSymbol({ kind }: { kind: Component }) {
  const file = kind.toLowerCase().replaceAll(" ", "-");
  return (
    <Image
      className="component-symbol"
      src={`/components/${file}.svg`}
      alt=""
      width={64}
      height={64}
      unoptimized
      loading="eager"
      draggable={false}
    />
  );
}
