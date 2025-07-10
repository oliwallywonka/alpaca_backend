package tourfx

import (
	"go.uber.org/fx"

	"github.com/labstack/echo/v4"
	/* "github.com/oliwallywonka/alpaca_backend/internal/tour/handlers"
	"github.com/oliwallywonka/alpaca_backend/internal/tour/services"
	"github.com/oliwallywonka/alpaca_backend/internal/tour/repositories" */
)

var Module = fx.Module(
	"tour",
	fx.Provide(
		NewHandler,
		NewService,
	),
	fx.Invoke(func(e *echo.Echo, h *Handler) {
		e.GET("/tours", h.GetAll)
		e.POST("/tours", h.Save)

		// REGISTER ROUTES
		/* e.POST("/tours/name_exists", h.NameExists)
		e.GET("/tours", h.GetPaginated)
		e.GET("/tours/:key", h.GetTourByUniqueKey)
		e.POST("/tours", h.SaveTour)
		e.PUT("/tours/:id", h.UpdateTour)

		e.GET("/tours/:id/images", h.GetTourImages)
		e.POST("/tours/:id/images", h.SetImage)
		e.PUT("/tours/:id/images", h.AddImage)
		e.DELETE("/tours/:id/images/:imageID", h.DeleteImage)

		e.GET("/tours/:tourID/destinations", h.GetTourDestinations)
		e.POST("/tours/:tourID/destinations", h.CreateTourDestination)
		e.PUT("/tours/:tourID/destinations/:tourDestinationID", h.UpdateTourDestination)
		e.DELETE("/tours/:tourID/destinations/:tourDestinationID", h.DeleteTourDestination)

		e.GET("/tours/:tourID/itineraries", h.GetItineraries)
		e.POST("/tours/:tourID/itineraries", h.CreateItinerary)
		e.PUT("/tours/:tourID/itineraries/:itineraryID", h.UpdateItinerary)
		e.DELETE("/tours/:tourID/itineraries/:itineraryID", h.DeleteItinerary)
		e.POST("/tours/:tourID/itineraries/:itineraryID/images", h.SetItineraryImage)
		e.DELETE("/tours/:tourID/itineraries/:itineraryID/images/:imageID", h.DeleteItineraryImage)

		e.GET("/tours/:tourID/meals", h.GetTourMeals)
		e.POST("/tours/:tourID/meals", h.CreateTourMeal)
		e.PUT("/tours/:tourID/meals/:tourMealID", h.UpdateTourMeal)
		e.DELETE("/tours/:tourID/meals/:tourMealID", h.DeleteTourMeal)

		e.GET("/tours/:tourID/activities", h.GetTourActivities)
		e.POST("/tours/:tourID/activities", h.CreateTourActivity)
		e.PUT("/tours/:tourID/activities/:tourActivityID", h.UpdateTourActivity)
		e.DELETE("/tours/:tourID/activities/:tourActivityID", h.DeleteTourActivity) */

		/*tour, err := r.GetByUniqueKey("123")
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		fmt.Printf("%+v\n", tour)*/
	}),
)
