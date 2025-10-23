package auction_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"fullcycle-auction_go/internal/entity/auction_entity"
	"fullcycle-auction_go/internal/infra/database/auction"
)

type CreateAuctionTestSuite struct {
	suite.Suite
	container  *mongodb.MongoDBContainer
	reposiotry *auction.AuctionRepository
}

func (s *CreateAuctionTestSuite) SetupSuite() {
	ctx := context.Background()
	container, err := mongodb.Run(ctx, "mongo:7")
	s.Require().NoError(err)
	connStr, err := container.ConnectionString(ctx)
	s.Require().NoError(err)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connStr))
	s.Require().NoError(err)

	s.reposiotry = auction.NewAuctionRepository(client.Database("testdb"))
	s.container = container
}

func (s *CreateAuctionTestSuite) TearDownSuite() {
	s.Require().NoError(s.container.Terminate(context.Background()))
}

func (s *CreateAuctionTestSuite) TestCreateAuction() {
	s.Run("when create auction then close automatically after interval duration", func() {
		ctx := context.Background()

		os.Setenv("AUCTION_INTERVAL", "200ms")
		defer os.Unsetenv("AUCTION_INTERVAL")

		auctionEntity := &auction_entity.Auction{
			Id:          "auction123",
			ProductName: "Test Product",
			Category:    "Test Category",
			Description: "Test Description",
			Condition:   auction_entity.New,
			Status:      auction_entity.Active,
			Timestamp:   time.Now(),
		}

		err := s.reposiotry.CreateAuction(ctx, auctionEntity)
		if err != nil {
			s.T().Fatalf("expected no error, got %v", err)
		}

		s.Equal(auction_entity.Active, auctionEntity.Status)

		// Aguarda o fechamento automático e a goroutine finalizar
		time.Sleep(300 * time.Millisecond)

		auction, err := s.reposiotry.FindAuctionById(ctx, auctionEntity.Id)
		if err != nil {
			s.T().Fatalf("expected no error on find, got %v", err)
		}

		s.Equal(auction_entity.Completed, auction.Status)
	})
}

func TestCreateAuctionTestSuite(t *testing.T) {
	suite.Run(t, new(CreateAuctionTestSuite))
}
