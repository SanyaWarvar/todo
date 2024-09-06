from faker import Faker
import aiohttp
import asyncio
import json
import random

fake = Faker(['en', 'ru_RU', 'ja_JP'])
BASE_URL = ""
class BearerAuth(aiohttp.BasicAuth):
    def __init__(self, token: str):
        self.token = token

    def encode(self) -> str:
        return f'Bearer: {self.token}'        
    

async def create_data(users_n: int, lists_n: int, tasks_n: int):
    err_num = 0
    one_percent = round(users_n * 0.01)
    for i in range(users_n):
        if i % one_percent == 0:
            print(f"complete: {int(i / users_n * 100)}%")

        async with aiohttp.ClientSession() as session:
            user = {"username": fake.name(), "password": fake.password()}  #имена
            async with session.post('http://localhost:80/auth/sign-up', data=json.dumps(user)) as resp:
                if resp.status != 201: #обычно из за дубликатов имени, тк библиотека не так уж и много уникальных данных генерирует
                    err_num += 1
                    continue
                pass

            async with session.post('http://localhost:80/auth/sign-in', data=json.dumps(user)) as resp:
                data = await resp.text()
                data = eval(data)

        async with aiohttp.ClientSession(auth=BearerAuth(data['token'])) as session:   
            for _ in range(lists_n):
                list = {"title": fake.text(random.randint(20, 40)), "description": fake.text(random.randint(20, 100))}
                async with session.post('http://localhost:80/api/lists', data=json.dumps(list)) as resp:
                    data = await resp.text()
                    data = eval(data)
                    if resp.status != 201:
                        err_num += 1
                        continue
                    list_id = data['list_id']
                
                for _ in range(tasks_n):
                    task = {"title": fake.text(random.randint(20, 30)), "description": fake.text(random.randint(20, 100))}
                    async with session.post(f'http://localhost:80/api/lists/{list_id}/items', data=json.dumps(task)) as resp:
                        pass

    print(f"Количество ошибок: {err_num}")   

async def main():
    users_n = 100
    list_n = 50
    tasks_n = 10
    await create_data(users_n, list_n, tasks_n)
    

if __name__ == "__main__":
       asyncio.run(main())