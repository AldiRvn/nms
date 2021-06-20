// Fill #list-database with value
new Vue({
    el: '#sidebar',
    data() {
        return {
            databases: []
        }
    },
    async mounted() {
        try {
            const databases = await axios.get(`/svc/database/collection/find`)
            // console.log(databases)

            databases.data.data.forEach(async database => {
                // console.log(database);

                this.databases.push({
                    name: database.name,
                    collections: database.collections
                })
            });
        } catch (err) {
            console.error(err)
        }
    }
})
