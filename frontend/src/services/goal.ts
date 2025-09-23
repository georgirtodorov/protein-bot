import { apiRequest } from "./api";

export const GoalService = {
  getGoal: () => apiRequest<{ goal: number }>("/v1/goal"),
  setGoal: (goal: number) =>
    apiRequest<{ success: boolean }>("/v1/goal", {
      method: "POST",
      body: JSON.stringify({ goal }),
    }),
};
