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


function callHostPortApi()
{
	//var obj  = '{"host":"'++'"}'
	alert("searching based on host port no. ...");
}


function callAppIdApi()
{
	alert("search based on appid..");
}