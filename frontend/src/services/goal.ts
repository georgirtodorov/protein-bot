import { ProteinGoalType } from "../types/ProteinGoalType";
import { apiRequest } from "./api";

export const GoalService = {
  getGoal: () => apiRequest<ProteinGoalType>("/v1/protein/goal"),
  setGoal: (amount: number) =>
    apiRequest<{ success: boolean }>("/v1/protein/goal", {
      method: "POST",
      body: JSON.stringify({ amount }),
    }),
  getHistory: () => apiRequest<ProteinGoalType[]>("/v1/protein/goal/history")
};
