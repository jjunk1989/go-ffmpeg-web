$(document).ready(function() {
	console.log("loaded")
	$("#check").on("click", function(e) {
		var s = Date.now()
		$.get("/api/test", 
		{
			type : "version"
		},
		function(res) {
			console.log("test connect server", res)
			if (!res) 
				return
			var cost = (Date.now() - s)
			if (res.code === 0) {
				$("#check").next().text("connect success " + res.result.version + ".cost： " + cost + "ms")
			} else {
				$("#check").next().text("connect failed" + res.message )
			}
		})
	})

	$("#testConver").on("click", function (e) {
		var s = Date.now()
		$.get("/api/test", 
		{
			type : "cmd"
		},
		function(res) {
			console.log("test conver", res)
			if (!res) 
				return
			var cost = (Date.now() - s)
			if (res.code === 0) {
				$("#testConver").next().text("conver failed " + res.result.version + ".cost： " + cost + "ms")
			} else {
				$("#testConver").next().text("conver failed" + res.message + ".cost： " + cost + "ms")
			}
		})
	})
	$("#getTask").on("click", function(e) {
		var s = Date.now()
		$.get("/api/task/" + $("#taskId").val(),{} ,function(res) {
			console.log("test conver", res)
			if (!res) 
				return
			var cost = (Date.now() - s)
			$("#getTask").next().text(res.message)
		})
	})
	$("#compose").on("click", function(e) {
		e.preventDefault()

		console.log("on compose")
		var s = Date.now()
		$("#compose").attr("disabled", "disabled")
		var formData = new FormData()

		if ($("#gif")[0].files.length > 0 && $("#mp3")[0].files.length > 0) {
			$("#gif").next().text("start")
			var gif = $("#gif")[0].files[0]
			var mp3 = $("#mp3")[0].files[0]
			formData.append("time", $("#time").val())
			formData.append("startTime", $("#startTime").val());
			formData.append("files", gif)
			formData.append("files", mp3)
			console.log("upload form data", formData)

            $.ajax({
                url: "/api/compose",
                type: "POST",
                data: formData,
                /**
                *必须false才会自动加上正确的Content-Type
                */
                contentType: false,
                /**
                * 必须false才会避开jQuery对 formdata 的默认处理
                * XMLHttpRequest会对 formdata 进行正确的处理
                */
                processData: false,
                success: function (data) {
					var cost = (Date.now() - s)
					$("#compose").removeAttr("disabled", "disabled")
                    if (data.code == 0) {
						$("#compose").next().text("upload success. start compose" + data.message)
                        checkTask(data.result.task)
                    } else {
                        $("#compose").next().text("upload failed" + data.message)
                    }
                },
                error: function (e) {
					console.log("upload failed:", e)
					$("#compose").removeAttr("disabled", "disabled")
                    $("#compose").next().text("upload failed:" + e)
                }
            }); 
		} else {
			if (!$("#gif")[0].files.length)
				$("#gif").next().text("need GIF PLZ!")
			if (!$("#mp3")[0].files.length)
				$("#mp3").next().text("need MP3 PLZ!")
		}
	})
	$("#gif").on("change", function(e) {
		console.log(e)
		$("#gif").next().text("got a gif")
	})
	$("#mp3").on("change", function(e) {
		console.log(e)
		$("#mp3").next().text("got a mp3")
	})
	function checkTask(tid) {
		$.get("/api/task/" + tid,{} ,function(res) {
			console.log("test conver", res)
			if (!res || res.code != 0)
				return;
			if (res.result.status == 1){
				$("#compose").next().text("compose success!")
				$("#composeVideo").attr("src", "/upload/" + res.result.output)
			} else if (res.result.status == 3) {
				$("#compose").next().text("compose failed!")	
			} else {
				window.setTimeout(function() {
					checkTask(tid)
				}, 1000)
			}
		})
	}
})