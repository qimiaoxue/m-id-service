from locust import HttpUser, task


class MonkeyDemo(HttpUser):
    @task
    def check_health(self):
        self.client.get("/vi/health")