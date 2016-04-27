
    //var endPoint = "http://localhost:8888";
    var endPoint = "http://dev--mapi--mapi--94ad2d.tx3.shipped-cisco.com";
    var appEndPoint = endPoint + "/app/{{appName}}";
    var hostPortEndpoint = endPoint + "/hostport/{{hostName}}/{{port}}";


    $(window).load(function() {

    });

    $(document).ready(function() {
            $('#s1').show();
        $("#btnSearch").click(function(){
        //alert("search based on appid..");
            if($("#host").val() != "")
            {
                if($("#port").val() != "")
                {
                    callHostPortApi();

                }
                else{
                    alert("Port no is missing")
                    }
            }
            else if($("#app").val() != "")
            {
                callAppIdApi();
            }
            else{
                alert("Search By either Host Name  & Port No or Search By App Id")
                }
        });
    });

    function renderApiData(data){
    if(typeof data === 'object'){

           $('.app-details-row').remove('.task-link');
        $('.app-details-row').remove('.task-content');
             $('#id').text(data.id);
               $('#project_name').text(data.projectname);
               $('#project_id').text(data.projectid);
               $('#env_name').text(data.envname);
               $('#env_id').text(data.envid);
               $('#service_name').text(data.servicename);
               $('#service_id').text(data.serviceid);

               $.each(data.tasks, function(i){
                var task = this;
                   var $view = $($('#tasks_list').html());
                   $view.find('.task-link').attr('href', '#task-'+i).text(task.id);
                   $view.find('.task-content').attr('id', 'task-'+i);
                   $view.find('.host').text(task.host);
                   $view.find('.id').text(task.id);
                   $view.find('.app-id').text(task.appid);
                   $view.find('.ports').text(task.ports.join());
                   $('.app-details-row').append($view);
                   //console.log($view.html());
               });

               $('.app-details').fadeIn(100);
           }
    }

    function callHostPortApi()
    {
        var hostName = $('#host').val().trim();
        var port = $('#port').val().trim();
        var requestURL = hostPortEndpoint.replace('{{hostName}}', encodeURI(hostName)).replace('{{port}}', encodeURI(port));

        $.get(requestURL, function(data){

          renderApiData(data);
        });


        //var obj  = '{"host":"'++'"}'
        //alert("searching based on host port no. ...");
    }


    function callAppIdApi()
    {
        var appName = $("#app").val().trim();
        var requestURL = appEndPoint.replace('{{appName}}', encodeURI(appName));
        $.get(requestURL, function(data){

          renderApiData(data);
        });

        //alert("search based on appid..");
    }