var app = new Vue({
    el: '#app',
    data: {
        recipes_objects: {},
        path: 'http://3.19.72.131:8000',
        modal_title: 'Update recipe',
        button_modal: 'save',
        id_object_selected: '',
        search:''
    },
    created: function () {
        this.load_table()
    },
    methods: {
        load_table: function () {
            $.ajax({
                url: this.path + "/allRepice",
                type: 'GET',
                crossDomain: true,
                success: function (result) {
                    app.recipes_objects = JSON.parse(result)
                },
                error: function (error) {
                }
            });
        },
        find_table: function () {
          let data_object = `{"name": "${app.search}"}`
          const settings = {
            "url": app.path + '/findRecipe',
            "method": "POST",
            "data": data_object,
            "crossDomain": true,
             
          }
          
          $.ajax(settings).done(function (response) {
            console.log(response)
            app.recipes_objects = JSON.parse(response);

          });
      },
        delete_recipe: function (id) {
            const id_object = app.recipes_objects[id].id
            const data_object = `{"ID": "${id_object}"}`
            Swal.fire({
                title: 'Are you sure?',
                text: "You won't be able to revert this!",
                type: 'warning',
                showCancelButton: true,
                confirmButtonColor: '#3085d6',
                cancelButtonColor: '#d33',
                confirmButtonText: 'Yes, delete it!'
              }).then((result) => {
                if (result.value) {
                  Swal.fire(
                    'Deleted!',
                    'Your recipe "'+app.recipes_objects[id].title +'" has been deleted.',
                    'success'
                  )
                  const settings = {
                    "url": app.path + '/deleteRecipe',
                    "method": "POST",
                    "data": data_object,
                    "crossDomain": true, 
                  }
                  
                  $.ajax(settings).done(function (response) {
                    console.log(response)
                    app.load_table()
    
                  });
                }
              })
      
        },
        update_recipe_load: function (id) {
            const id_object = app.recipes_objects[id]
            app.modal_title = 'Update recipe'
            app.button_modal = 'Update'
            app.id_object_selected = app.recipes_objects[id].id
            $('#name').val(id_object.title)
            $('#ingredients').val(id_object.ingredients)
            $('#preparation').val(id_object.preparation)
            
        },
        add_recipe_load: function () {

          app.modal_title = 'Add new recipe'
          app.button_modal = 'Create'
          $('#name').val('')
          $('#ingredients').val('')
          $('#preparation').val('')
          
      },
      save_recipe: function (){
        let ruta = '/insertRecipe'
        let type = 'created'
        let data_send = `{"Title":"${$('#name').val().replace(/(\r\n|\n|\r)/gm," ")}","Ingredients":"${$('#ingredients').val().replace(/(\r\n|\n|\r)/gm," ")}","Preparation":"${$('#preparation').val().replace(/(\r\n|\n|\r)/gm," ")}"}`
        if(app.modal_title == 'Update recipe'){
          ruta = '/updateRecipe'
          type = 'updated'
          data_send = `{"Id": "${app.id_object_selected}","Title":"${$('#name').val().replace(/(\r\n|\n|\r)/gm," ")}","Ingredients":"${$('#ingredients').val().replace(/(\r\n|\n|\r)/gm," ")}","Preparation":"${$('#preparation').val().replace(/(\r\n|\n|\r)/gm," ")}"}`
        }
        var settings = {
          "url": app.path + ruta,
          "method": "POST",
          "data": data_send,
          "crossDomain": true,
        }
        
        $.ajax(settings).done(function (response) {
          console.log(response);
          Swal.fire(
            'Good job!',
            'Your recipe has been ' + type + '!',
            'success'
          )
          app.load_table()
        });
      }
        
    }
});
