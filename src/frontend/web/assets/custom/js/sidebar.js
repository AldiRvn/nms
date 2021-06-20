// Fill #list-database with value
const SIDE_BAR = {
    data() {
        return {
            databases: [
                {
                    name: "mock",
                    collections: [
                        "users",
                        "books"
                    ]
                },
                {
                    name: "kiasan",
                    collections: [
                        "semua",
                        "abadi",
                        "kemari"
                    ]
                }
            ]
        }
    }
}

Vue.createApp(SIDE_BAR).mount("#sidebar")