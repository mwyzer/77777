import { useQuery } from "@tanstack/react-query";
import { Link } from "react-router-dom";
import { getCustomers } from "../services/api";
import type { Customer } from "../types";

export default function CustomersPage() {
  const { data, isLoading, error } = useQuery({
    queryKey: ["customers"],
    queryFn: () => getCustomers(),
  });

  return (
    <div>
      <h2 className="text-xl font-semibold mb-4">Customers</h2>

      {isLoading && (
        <div className="text-center py-8 text-gray-500">
          Loading customers...
        </div>
      )}

      {error && (
        <div className="p-4 bg-red-50 border border-red-200 text-red-700 rounded">
          Failed to load customers.
        </div>
      )}

      {data && data.customers.length === 0 && (
        <div className="text-center py-8 text-gray-400">No customers yet.</div>
      )}

      {data && data.customers.length > 0 && (
        <div className="bg-white rounded-lg border border-gray-200 overflow-hidden">
          <table className="w-full text-sm">
            <thead className="bg-gray-50 border-b">
              <tr>
                <th className="text-left px-4 py-3 font-medium text-gray-600">
                  Name
                </th>
                <th className="text-left px-4 py-3 font-medium text-gray-600">
                  Phone
                </th>
                <th className="text-left px-4 py-3 font-medium text-gray-600">
                  Provider
                </th>
                <th className="text-left px-4 py-3 font-medium text-gray-600">
                  Joined
                </th>
              </tr>
            </thead>
            <tbody>
              {data.customers.map((c: Customer) => (
                <tr key={c.id} className="border-b last:border-0">
                  <td className="px-4 py-3">
                    <Link
                      to={`/customers/${c.id}`}
                      className="text-blue-600 hover:underline font-medium"
                    >
                      {c.name}
                    </Link>
                  </td>
                  <td className="px-4 py-3 text-gray-600">{c.phone || "-"}</td>
                  <td className="px-4 py-3">
                    <span
                      className={`text-xs px-2 py-0.5 rounded-full ${
                        c.provider === "whatsapp"
                          ? "bg-green-100 text-green-700"
                          : "bg-blue-100 text-blue-700"
                      }`}
                    >
                      {c.provider}
                    </span>
                  </td>
                  <td className="px-4 py-3 text-gray-500">
                    {new Date(c.created_at).toLocaleDateString()}
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}
    </div>
  );
}
