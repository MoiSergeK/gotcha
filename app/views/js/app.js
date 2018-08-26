var socket = null;


$(document).ready(() => {
    initSocket();

    loadSelfPlaces();
    loadCommonPlaces();
    loadUsers();
    
    showLog();
    
    setInterval(() => {
        showLog();
    }, 5000);

    $('#addCommonPlaceBtn').on('click', (e) => {
        $('#addCommonPlaceModal').modal('show');
    })
})

function showLog(){
    $.get('/log', (response) => {
        $('#logStack').html("");

        for(let line of response.data.split("\n")){
            $('#logStack').append('<small>' + line + '</small><br>');
        }
    })
}

function clearLog(){
    $.ajax({url: '/log', method: 'DELETE', success: (response) => {
        showLog();
    }})
}

function loadUsers(){
    $.get('/users', (response) => {
        $('#usersHolder').html('');

        let rows = response.data;

        if(rows){
            for(let row of rows) {
                let tr = $('<tr></tr>');

                tr.append('<td>' + row.id + '</td>');
                tr.append('<td>' + row.name + '</td>');
                tr.append('<td>' + row.phone + '</td>');

                $('#usersHolder').append(tr)
            }
        }
    })
}

function loadSelfPlaces(){
    $.get('/places/self', (response) => {
        $('#myPlacesHolder').html('');

        let rows = response.data;

        if(rows){
            for(let row of rows) {
                let tr = $('<tr></tr>');

                tr.append('<td>' + row.id + '</td>');
                tr.append('<td>' + row.lat + ", " + row.lng + '</td>');
                tr.append('<td>' + row.address + '</td>');

                let td = $('<td></td>');

                let btn = $('<button class="btn btn-outline-danger btn-sm float-right">Del</button>');

                btn.on('click', (e) => {
                    $.ajax({url: '/places/self/' + row.id, method: 'DELETE', success: (response) => {
                        loadSelfPlaces();
                    }})
                })

                td.append(btn);

                tr.append(td);

                $('#myPlacesHolder').append(tr)
            }
        }
    })
}

function loadCommonPlaces(){
    $.get('/places/common', (response) => {
        $('#commonPlacesHolder').html('');

        let rows = response.data;

        if(rows){
            for(let row of rows) {
                let tr = $('<tr></tr>');

                tr.append('<td>' + row.id + '</td>');
                tr.append('<td>' + row.lat + ", " + row.lng + '</td>');
                tr.append('<td>' + row.name + '</td>');

                let td = $('<td></td>');

                let btn = $('<button class="btn btn-outline-danger btn-sm float-right">Del</button>');

                btn.on('click', (e) => {
                    $.ajax({url: '/places/common/' + row.id, method: 'DELETE', success: (response) => {
                        loadCommonPlaces();
                    }})
                })

                td.append(btn);

                tr.append(td);

                $('#commonPlacesHolder').append(tr)
            }
        }
    })
}

function addCommonPlace(){
    let coords = $('#addCommonPlaceForm').find('input[name=coords]').val().split(',').map(x => x.trim())
    let sendData = {
        name: $('#addCommonPlaceForm').find('input[name=name]').val().trim(),
        lat: coords[0],
        lng: coords[1],
    };

    $.post('/places/common', JSON.stringify(sendData), (resp) => {
        $('#addCommonPlaceModal').modal('hide');
        loadCommonPlaces();
    })
}

function searchNearby(){
    let lat = $('#searchNearbyForm').find('input[name=lat]').val().trim();
    let lng = $('#searchNearbyForm').find('input[name=lng]').val().trim();

    $.get('/places/nearby', {lat: lat, lng: lng}, (response) => {
        $('#nearbyPlaces').html('')

        for(let p of response.data) {
            $('#nearbyPlaces').append(p.name + '<br>')
        }
    })
}

function initSocket(){
    // socket = io.connect('http://127.0.0.1:8080');

    // socket.on('connect', function(){
    //     console.log('connect');

    //     socket.emit('event', 'hello');
    // });
}