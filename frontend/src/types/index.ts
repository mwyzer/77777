// ── Auth ──
export interface User {
  id: string;
  name: string;
  email: string;
  role: string;
  created_at: string;
  updated_at: string;
}

export interface LoginRequest {
  email: string;
  password: string;
}

export interface LoginResponse {
  token: string;
  user: User;
}

// ── Inbox ──
export interface Customer {
  id: string;
  name: string;
  phone: string;
  email: string;
  provider: string;
  provider_id: string;
  created_at: string;
  updated_at: string;
}

export interface Conversation {
  id: string;
  customer_id: string;
  customer?: Customer;
  channel: string;
  status: string;
  last_message_at: string | null;
  created_at: string;
  updated_at: string;
}

export interface Message {
  id: string;
  conversation_id: string;
  sender_type: "customer" | "agent";
  content: string;
  status: string;
  provider_message_id: string | null;
  created_at: string;
}

// ── API Responses ──
export interface APIResponse<T = unknown> {
  success: boolean;
  message?: string;
  data?: T;
  error?: string;
}

export interface PaginatedList<T> {
  [key: string]: T[] | number;
  total: number;
  page: number;
  limit: number;
}

export interface ConversationListResponse extends PaginatedList<Conversation> {
  conversations: Conversation[];
}

export interface MessageListResponse extends PaginatedList<Message> {
  messages: Message[];
}

export interface CustomerListResponse extends PaginatedList<Customer> {
  customers: Customer[];
}

export interface CustomerDetailResponse {
  customer: Customer;
  conversations: Conversation[];
}
