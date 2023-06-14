import psycopg2
from faker import Faker

# username := postgres
# password := postgres
# host := localhost
# port := 5432
# dbname := test
# docker-container := psql-dev

conn = psycopg2.connect(
    host="localhost",
    database="test",
    user="postgres",
    password="postgres"
)
cur = conn.cursor()

fake = Faker()

for i in range(50):
    movie_title = fake.text(max_nb_chars=50)
    release_year = fake.random_int(min=1900, max=2023)
    director = fake.name()
    genre = fake.word()
    cur.execute("INSERT INTO Movies (title, release_year, director, genre) VALUES (%s, %s, %s, %s)", (movie_title, release_year, director, genre))


for i in range(100):
    actor_name = fake.name()
    birth_year = fake.random_int(min=1900, max=2023)
    nationality = fake.country()
    cur.execute("INSERT INTO Actors (name, birth_year, nationality) VALUES (%s, %s, %s)", (actor_name, birth_year, nationality))

awards_data = [
    ("Best Picture", "Motion Picture"),
    ("Best Actor", "Lead Actor"),
    ("Best Actress", "Lead Actress"),
    ("Best Supporting Actor", "Supporting Actor"),
    ("Best Supporting Actress", "Supporting Actress")
]
for award_name, category in awards_data:
    cur.execute("INSERT INTO awards (name, category) VALUES (%s, %s)", (award_name, category))

for i in range(100):
    movie_id = fake.random_int(min=1, max=50)
    award_id = fake.random_int(min=1, max=5)
    year = fake.random_int(min=1900, max=2023)
    is_winner = fake.boolean(chance_of_getting_true=50)
    cur.execute("INSERT INTO nominations (movie_id, award_id, year, is_winner) VALUES (%s, %s, %s, %s)", (movie_id, award_id, year, is_winner))


for i in range(200):
    actor_id = fake.random_int(min=1, max=100)
    movie_id = fake.random_int(min=1, max=50)
    year = fake.random_int(min=1900, max=2023)
    cur.execute("INSERT INTO performances (actor_id, movie_id, year) VALUES (%s, %s, %s)", (actor_id, movie_id, year))


nominations_data = cur.execute("SELECT id FROM nominations")
nominations = [n[0] for n in cur.fetchall()]
performances_data = cur.execute("SELECT id FROM performances")
performances = [p[0] for p in cur.fetchall()]
for i in range(100):
    nomination_id = fake.random_element(elements=nominations)
    performance_id = fake.random_element(elements=performances)
    try:
      cur.execute("INSERT INTO Nominated_Performances (nomination_id, performance_id) VALUES (%s, %s)", (nomination_id, performance_id))
    except:
      pass

conn.commit()
cur.close()
conn.close()