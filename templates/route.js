var btn = document.getElementById('submit_f');

btn.onclick = function() {


	
	
var text_name=document.getElementById("text_name").value;
var text_username=document.getElementById("text_username").value;

	alert(text_name);alert(text_username);
	
	
	jQuery.ajax({
	url: "",
	data: {'value':text_name,'key':text_username},
	type: "POST",
	success:function(dataa){
		if(dataa==true)
		{
	alert("Submit Successful")
		}

	}
	
	});
	
}



var btn = document.getElementById('submit_s');

btn.onclick = function() {


	
	
var text_username_s=document.getElementById("text_username_s").value;


	alert(text_username_s);
	
	
	jQuery.ajax({
	url: "",
	data: {'key':text_username_s},
	type: "POST",
	success:function(dataa){
		if(dataa==true)
		{
	alert("Subscribe Successful")
		}

	}
	
	});
	
}