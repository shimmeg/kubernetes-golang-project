db.createUser({
    user: 'aefimov',
    pwd: 'aefimov123',
    roles: [
        {
            role: 'readWrite',
            db: 'tasks-tracker'
        }
    ]
})
db = new Mongo().getDB("tasks-tracker");

db.createCollection('tasks', { capped: false });