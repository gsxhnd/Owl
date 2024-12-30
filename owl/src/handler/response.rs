use axum::{http::StatusCode, Json};
use serde::Serialize;

#[derive(Debug, Clone, Serialize)]
pub struct HandlerResponse<T> {
    data: Option<T>,
    message: String,
}

impl<T> Default for HandlerResponse<T> {
    fn default() -> Self {
        Self {
            data: None,
            message: String::default(),
        }
    }
}

impl<T> HandlerResponse<T> {
    pub fn new(data: Option<T>, message: impl ToString) -> Json<Self> {
        let message = message.to_string();
        Json(Self { data, message })
    }

    pub fn ok(content: T) -> (StatusCode, Json<Self>) {
        (StatusCode::OK, Self::new(Some(content), "OK"))
    }

    pub fn err(message: impl ToString) -> (StatusCode, Json<Self>) {
        (StatusCode::INTERNAL_SERVER_ERROR, Self::new(None, message))
    }

    pub fn not_found(message: impl ToString) -> (StatusCode, Json<Self>) {
        (StatusCode::NOT_FOUND, Self::new(None, message))
    }

    pub fn internal_server_error() -> (StatusCode, Json<Self>) {
        (
            StatusCode::INTERNAL_SERVER_ERROR,
            Self::new(None, "Internal server error"),
        )
    }
}
