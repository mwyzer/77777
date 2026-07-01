import { useQuery } from "@tanstack/react-query";
import { Link } from "react-router-dom";
import { getConversations } from "../services/api";
import type { Conversation } from "../types";

export default function InboxPage() {
  const { data, isLoading, error } = useQuery({
    queryKey: ["conversations"],
    queryFn: () => getConversations(),
  });

  return (
    <div>
      <h2 className="text-xl font-semibold mb-4">Inbox</h2>

      {isLoading && (
        <div className="text-center py-8 text-gray-500">
          Loading conversations...
        </div>
      )}

      {error && (
        <div className="p-4 bg-red-50 border border-red-200 text-red-700 rounded">
          Failed to load conversations.
        </div>
      )}

      {data && data.conversations.length === 0 && (
        <div className="text-center py-8 text-gray-400">
          No conversations yet. Messages from WhatsApp or Telegram will appear
          here.
        </div>
      )}

      {data && data.conversations.length > 0 && (
        <div className="space-y-2">
          {data.conversations.map((conv: Conversation) => (
            <Link
              key={conv.id}
              to={`/conversations/${conv.id}`}
              className="block p-4 bg-white rounded-lg border border-gray-200 hover:border-blue-300 hover:shadow-sm transition"
            >
              <div className="flex items-center justify-between">
                <div>
                  <span className="font-medium text-gray-900">
                    {conv.customer?.name || "Unknown"}
                  </span>
                  <span className="ml-2 text-xs text-gray-500">
                    {conv.customer?.phone && `· ${conv.customer.phone}`}
                  </span>
                </div>
                <div className="flex items-center gap-2">
                  <span
                    className={`text-xs px-2 py-0.5 rounded-full ${
                      conv.channel === "whatsapp"
                        ? "bg-green-100 text-green-700"
                        : "bg-blue-100 text-blue-700"
                    }`}
                  >
                    {conv.channel}
                  </span>
                  {conv.status === "open" && (
                    <span
                      className="w-2 h-2 rounded-full bg-green-500"
                      title="Open"
                    />
                  )}
                </div>
              </div>
              {conv.last_message_at && (
                <p className="text-xs text-gray-400 mt-1">
                  Last message:{" "}
                  {new Date(conv.last_message_at).toLocaleString()}
                </p>
              )}
            </Link>
          ))}
        </div>
      )}
    </div>
  );
}
