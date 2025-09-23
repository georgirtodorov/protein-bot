import { apiRequest } from "./api";

export interface ProteinEntry {
  id?: number;
  amount: number;
  description?: string;
  created_at?: string;
}

export interface AddProteinResponse {
  added: number;
  total: number;
  goal: number;
  remaining: number;
}

export const ProteinService = {
  add: (entry: Omit<ProteinEntry, "id" | "created_at">) =>
    apiRequest<AddProteinResponse>("/v1/add", {
      method: "POST",
      body: JSON.stringify(entry),
    }),

  history: () => apiRequest<ProteinEntry[]>("/v1/history"),

  status: () => apiRequest<{ total: number; goal: number }>("/v1/status"),
};
