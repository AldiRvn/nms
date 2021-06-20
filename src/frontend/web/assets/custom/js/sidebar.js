// Fill #list-database with value
new Vue({
    el: '#sidebar',
    data () {
      return {
        databases: []
      }
    },
    async mounted () {
        try {
            const users = await axios.get(`https://jsonplaceholder.typicode.com/users`)
            // console.log(users)

            users.data.forEach(async user => {
                // console.log(user);

                const todos = []
                const user_todos = await axios.get(`https://jsonplaceholder.typicode.com/users/${user.id}/todos`)
                // console.log(user_todos)

                user_todos.data.forEach(todo => {
                    todos.push(todo.title)
                })

                this.databases.push({
                    name: user.name,
                    collections: todos
                })
            });
        } catch (err) {
            console.error(err)
        }
    }
  })
  