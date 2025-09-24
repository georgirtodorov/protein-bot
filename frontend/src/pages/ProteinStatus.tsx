import { useEffect, useState } from "react";
import { ProteinService } from "../services/protein";
import { ProteinStatusType } from "../types/ProteinStatusType";

export default function ProteinStatus({ refreshKey }: { refreshKey: number }) {
  const [status, setStatus] = useState<ProteinStatusType | null>(null);

  useEffect(() => {
    async function fetchStatus() {
      try {
        const data = await ProteinService.status();
        setStatus(data);
      } catch (err) {
        console.error("Failed to load status:", err);
      }
    }
    fetchStatus();
  }, [refreshKey]);

  if (!status) {
    return <div className="p-3 bg-gray-100 rounded-xl">Loading status...</div>;
  }

  return (
    <div className="p-3 bg-gray-100 rounded-xl mb-4">
      <p><strong>Total today:</strong> {status.total}g</p>
      <p><strong>Goal:</strong> {status.goal}g</p>
      <p><strong>Remaining:</strong> {status.remaining}g</p>
    </div>
  );
}
