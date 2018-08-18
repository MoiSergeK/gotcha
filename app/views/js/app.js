$(document).ready(() => {
    loadPlaces();
})

function loadPlaces(){
    $.get(location.protocol + '//' + location.host + '/places', (response) => {
        $('tbody').html('');

        let rows = response.data;

        if(rows){
            for(let row of rows) {
                let tr = $('<tr></tr>');

                for(let k in row)
                    tr.append('<td>' + row[k] + '</td>');

                let td = $('<td></td>');

                let btn = $('<button class="btn btn-outline-danger btn-sm ">Del</button>');

                btn.on('click', (e) => {
                    $.ajax({url: location.protocol + '//' + location.host + '/places/' + row.id, method: 'DELETE', success: (response) => {
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