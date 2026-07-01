import { useState } from "react";
import { useParams } from "react-router-dom";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { getConversation, getMessages, sendMessage } from "../services/api";
import type { Message } from "../types";

export default function ConversationDetailPage() {
  const { id } = useParams<{ id: string }>();
  const queryClient = useQueryClient();
  const [reply, setReply] = useState("");

  const { data: convData, isLoading: convLoading } = useQuery({
    queryKey: ["conversation", id],
    queryFn: () => getConversation(id!),
    enabled: !!id,
  });

  const { data: msgData, isLoading: msgLoading } = useQuery({
    queryKey: ["messages", id],
    queryFn: () => getMessages(id!),
    enabled: !!id,
  });

  const sendMutation = useMutation({
    mutationFn: (content: string) => sendMessage(id!, content),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["messages", id] });
      setReply("");
    },
  });

  const handleSend = (e: React.FormEvent) => {
    e.preventDefault();
    if (reply.trim()) {
      sendMutation.mutate(reply.trim());
    }
  };

  if (convLoading || msgLoading) {
    return <div className="text-center py-8 text-gray-500">Loading...</div>;
  }

  const customer = convData?.customer;

  return (
    <div className="flex flex-col h-[calc(100vh-8rem)]">
      {/* Header */}
      <div className="flex items-center gap-3 pb-4 border-b mb-4">
        <a href="/" className="text-blue-600 hover:underline text-sm">
          &larr; Back
        </a>
        <div>
          <h3 className="font-semibold">{customer?.name || "Unknown"}</h3>
          <p className="text-xs text-gray-500">
            {customer?.phone} · {convData?.conversation?.channel}
          </p>
        </div>
      </div>

      {/* Messages */}
      <div className="flex-1 overflow-y-auto space-y-3 mb-4">
        {msgData?.messages.map((msg: Message) => (
          <div
            key={msg.id}
            className={`flex ${msg.sender_type === "agent" ? "justify-end" : "justify-start"}`}
          >
            <div
              className={`max-w-[70%] px-4 py-2 rounded-lg text-sm ${
                msg.sender_type === "agent"
                  ? "bg-blue-600 text-white rounded-br-none"
                  : "bg-gray-100 text-gray-900 rounded-bl-none"
              }`}
            >
              <p>{msg.content}</p>
              <p
                className={`text-xs mt-1 ${msg.sender_type === "agent" ? "text-blue-200" : "text-gray-400"}`}
              >
                {new Date(msg.created_at).toLocaleTimeString()}
              </p>
            </div>
          </div>
        ))}

        {msgData?.messages.length === 0 && (
          <div className="text-center py-8 text-gray-400">No messages yet.</div>
        )}
      </div>

      {/* Reply box */}
      <form onSubmit={handleSend} className="flex gap-2 pt-4 border-t">
        <input
          type="text"
          value={reply}
          onChange={(e) => setReply(e.target.value)}
          placeholder="Type a reply..."
          className="flex-1 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
        />
        <button
          type="submit"
          disabled={!reply.trim() || sendMutation.isPending}
          className="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 disabled:opacity-50"
        >
          {sendMutation.isPending ? "Sending..." : "Send"}
        </button>
      </form>
    </div>
  );
}
