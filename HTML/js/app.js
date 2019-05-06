var app = new Vue({
    el: '#app',
    data: {
        recipes_objects: {},
        path: 'http://localhost:8000'
    },
    created: function () {
        this.load_table()
    },
    methods: {
        load_table: function () {
            $.ajax({
                url: this.path + "/allRepice",
                type: 'GET',
                success: function (result) {
                    app.recipes_objects = JSON.parse(result);
                    console.log("::: ", $.parseJSON(result))
                },
                error: function (error) {
                }
            });
        },
        delete_recipe: function (id) {
            console.log(">>> id <<< ", id)
            const id_object = app.recipes_objects[id].id
            console.log(">>> id_object <<< ", id_object)
            const data_object = `{\"ID\": ${id_object}}`
            var settings = {
                "url": "http://localhost:8000/deleteRecipe",
                "method": "POST",
                "data": data_object 
              }
              
              $.ajax(settings).done(function (response) {
                console.log(response);
              });
            /*
            console.log(">>> id <<< ", id)
            const id_object = app.recipes_objects[id].id
            const data_object = {"id": id_object }

            $.ajax({
                url: 'http://localhost:8000/deleteRecipe',
                type: 'POST',
                data: data_object,
                success: function (result) {
                    console.log(">>> result >>> ", result);
                    app.load_table()
                    
                },
                error: function (error) {
                    console.log("error<< ", error);
                }
            });*/
        },
    }
});
