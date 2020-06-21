package api

import (
        "net/http"
        "io/ioutil"
        "github.com/labstack/echo"
)

func NumbersHandler(ctx echo.Context) error {
        number := ctx.Param("number")

        url := "http://numbersapi.com/" + number + "?json"

        req, _ := http.NewRequest("GET", url, nil)
        resp, _ := http.DefaultClient.Do(req)

        defer resp.Body.Close()

        body, _ := ioutil.ReadAll(resp.Body)
        return ctx.JSON(http.StatusOK, string(body))
}
