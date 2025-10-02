import React, { ComponentProps } from "react";
import { Ionicons } from "@expo/vector-icons";

interface TabBarIconProps {
  name: ComponentProps<typeof Ionicons>["name"];
  color?: string;
  size?: number;
}

export default function TabBarIcon({ name, color = "black", size = 24 }: TabBarIconProps) {
  return <Ionicons name={name} size={size} color={color} />;
}