$(document).ready(() => {
    loadPlaces();
    
    showLog();
    
    setInterval(() => {
        showLog();
    }, 5000);
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

function loadPlaces(){
    $.get('/places', (response) => {
        $('tbody').html('');

        let rows = response.data;

        if(rows){
            for(let row of rows) {
                let tr = $('<tr></tr>');

                tr.append('<td>' + row.id + '</td>');
                tr.append('<td>' + row.coords + '</td>');
                tr.append('<td>' + row.address + '</td>');

                let td = $('<td></td>');

                let btn = $('<button class="btn btn-outline-danger btn-sm ">Del</button>');

                btn.on('click', (e) => {
                    $.ajax({url: '/places/' + row.id, method: 'DELETE', success: (response) => {
                        loadPlaces();
                    }})
                })

                td.append(btn);

                tr.append(td);

                $('tbody').append(tr)
            }
        }
    })
}