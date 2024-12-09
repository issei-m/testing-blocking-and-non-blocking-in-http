use actix_web::{get, web, App, HttpResponse, HttpServer, Responder};
use actix_web::middleware::Logger;
use env_logger::Env;
use serde::Deserialize;

#[derive(Deserialize)]
struct IndexParams {
    sleep: Option<String>,
}

#[get("/")]
async fn index(p: web::Query<IndexParams>) -> impl Responder {
    let sleep_time = p.sleep
        .as_deref()
        .and_then(|s| s.parse::<u64>().ok())
        .unwrap_or(10);

    tokio::time::sleep(std::time::Duration::from_secs(sleep_time)).await;

    let body_text = format!("{} seconds have passed", sleep_time);

    HttpResponse::Ok().body(body_text)
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    env_logger::init_from_env(Env::default().default_filter_or("info"));

    HttpServer::new(|| {
        App::new()
            .wrap(Logger::default())
            .service(index)
    })
        .bind(("0.0.0.0", 3000))?
        .run()
        .await
}
