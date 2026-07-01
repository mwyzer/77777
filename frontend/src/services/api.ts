import axios from "axios";
import { useAuthStore } from "../stores/authStore";
import type {
  LoginRequest,
  LoginResponse,
  ConversationListResponse,
  MessageListResponse,
  CustomerListResponse,
  CustomerDetailResponse,
  Message,
  User,
} from "../types";

const api = axios.create({
  baseURL: import.meta.env.VITE_API_URL || "/api",
  headers: { "Content-Type": "application/json" },
});

// Attach JWT token to every request
api.interceptors.request.use((config) => {
  const token = useAuthStore.getState().token;
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

// Handle 401 globally
api.interceptors.response.use(
  (res) => res,
  (err) => {
    if (err.response?.status === 401) {
      useAuthStore.getState().logout();
    }
    return Promise.reject(err);
  },
);

// ── Auth ──
export const login = (data: LoginRequest) =>
  api
    .post<{ success: boolean; data: LoginResponse }>("/auth/login", data)
    .then((r) => r.data.data);

export const getMe = () =>
  api
    .get<{ success: boolean; data: { user: User } }>("/auth/me")
    .then((r) => r.data.data.user);

// ── Conversations ──
export const getConversations = (params?: Record<string, string | number>) =>
  api
    .get<{
      success: boolean;
      data: ConversationListResponse;
    }>("/inbox/conversations", { params })
    .then((r) => r.data.data);

export const getConversation = (id: string) =>
  api
    .get<{
      success: boolean;
      data: {
        conversation: ConversationListResponse["conversations"][0];
        customer: CustomerListResponse["customers"][0];
      };
    }>(`/inbox/conversations/${id}`)
    .then((r) => r.data.data);

// ── Messages ──
export const getMessages = (
  conversationId: string,
  params?: Record<string, string | number>,
) =>
  api
    .get<{
      success: boolean;
      data: MessageListResponse;
    }>(`/inbox/conversations/${conversationId}/messages`, { params })
    .then((r) => r.data.data);

export const sendMessage = (conversationId: string, content: string) =>
  api
    .post<{
      success: boolean;
      data: Message;
    }>(`/inbox/conversations/${conversationId}/messages`, { content })
    .then((r) => r.data.data);

// ── Customers ──
export const getCustomers = (params?: Record<string, string | number>) =>
  api
    .get<{
      success: boolean;
      data: CustomerListResponse;
    }>("/inbox/customers", { params })
    .then((r) => r.data.data);

export const getCustomer = (id: string) =>
  api
    .get<{
      success: boolean;
      data: CustomerDetailResponse;
    }>(`/inbox/customers/${id}`)
    .then((r) => r.data.data);

export default api;
