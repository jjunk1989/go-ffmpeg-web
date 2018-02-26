$(document).ready(function() {
	console.log("loaded")
	$("#check").on("click", function(e) {
		var s = Date.now()
		$.get("/api/test", 
		{
			type : "version"
		},
		function(res) {
			console.log("连接服务器", res)
			if (!res) 
				return
			var cost = (Date.now() - s)
			if (res.code === 0) {
				$("#check").next().text("连接成功 " + res.result.version + ".耗时： " + cost + "ms")
			} else {
				$("#check").next().text("连接服务器失败" + res.message )
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
			console.log("测试转换", res)
			if (!res) 
				return
			var cost = (Date.now() - s)
			if (res.code === 0) {
				$("#testConver").next().text("转换成功 " + res.result.version + ".耗时： " + cost + "ms")
			} else {
				$("#testConver").next().text("转换失败" + res.message + ".耗时： " + cost + "ms")
			}
		})
	})
	
	$("#compose").on("click", function(e) {
		e.preventDefault()

		console.log("on compose")
		var s = Date.now()
		$("#compose").attr("disabled", "disabled")
		var formData = new FormData();

		if ($("#gif")[0].files.length > 0 && $("#mp3")[0].files.length > 0) {
			$("#gif").next().text("开始上传")
			var gif = $("#gif")[0].files[0]
			var mp3 = $("#mp3")[0].files[0]
			formData.append("files", gif);  
			formData.append("files", mp3);
			console.log("upload form data", formData);

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
                        $("#compose").next().text("上传成功! " + cost + "ms")
						$("#composeVideo").attr("src", "/upload/" + data.result.video)
                    } else {
                        $("#compose").next().text("上传失败" + data.message)
                    }
                },
                error: function (e) {
					console.log("上传失败:", e)
					$("#compose").removeAttr("disabled", "disabled")
                    $("#compose").next().text("上传失败:" + e)
                }
            }); 
		} else {
			if (!$("#gif")[0].files.length)
				$("#gif").next().text("请选择 GIF 文件!")
			if (!$("#mp3")[0].files.length)
				$("#mp3").next().text("请选择 MP3 文件!")
		}
	})
})